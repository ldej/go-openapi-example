package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	chimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/ldej/go-openapi-example/helpers"
	"github.com/ldej/go-openapi-example/oapi-codegen/api"
)

func main() {
	// Create your usual storage/service/handler layers
	store := helpers.NewThingsDatastore()
	service := NewThingService(store)
	s := NewServer(service)

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	// Add swagger UI endpoints
	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(swagger)
	})
	router.Handle("/swagger/", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		Path:    "/swagger/",
		SpecURL: "/swagger/doc.json",
	}, nil))

	// Enable validation of incoming requests
	validator := chimiddleware.OapiRequestValidatorWithOptions(
		swagger,
		&chimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
					return nil
				},
			},
		},
	)

	apiServer := api.HandlerWithOptions(
		api.NewStrictHandler(s, nil),
		api.ChiServerOptions{
			BaseURL:     "/api/v1",
			BaseRouter:  router,
			Middlewares: []api.MiddlewareFunc{validator},
		},
	)

	addr := ":8000"
	httpServer := http.Server{
		Addr:    addr,
		Handler: apiServer,
	}

	log.Println("Server listening on", addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewServer(thingService *ThingService) api.StrictServerInterface {
	return &server{
		ThingService: thingService,
	}
}

type server struct {
	*ThingService
}

type ThingService struct {
	store *helpers.ThingsDatastore
}

func NewThingService(store *helpers.ThingsDatastore) *ThingService {
	return &ThingService{
		store: store,
	}
}

func (s *ThingService) CreateThing(ctx context.Context, request api.CreateThingRequestObject) (api.CreateThingResponseObject, error) {
	thing := helpers.Thing{
		UUID:    uuid.NewString(),
		Name:    request.Body.Name,
		Rank:    request.Body.Rank,
		Rating:  helpers.ToFloat32(request.Body.Rating),
		Score:   helpers.ToFloat64(request.Body.Score),
		Type:    string(request.Body.Type),
		Created: time.Now().UTC(),
	}
	err := s.store.StoreThing(thing)
	if err != nil {
		return api.CreateThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.CreateThing200JSONResponse(mapThingToThingResponse(thing)), nil
}

func (s *ThingService) DeleteThing(ctx context.Context, request api.DeleteThingRequestObject) (api.DeleteThingResponseObject, error) {
	err := s.store.DeleteThing(request.Uuid.String())
	if err == helpers.ErrNotFound {
		return api.DeleteThing404JSONResponse{}, nil
	}
	if err != nil {
		return api.DeleteThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.DeleteThing204Response{}, nil
}

func (s *ThingService) GetThing(ctx context.Context, request api.GetThingRequestObject) (api.GetThingResponseObject, error) {
	thing, err := s.store.GetThing(request.Uuid.String())
	if err == helpers.ErrNotFound {
		return api.GetThing404JSONResponse{}, nil
	}
	if err != nil {
		return api.GetThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.GetThing200JSONResponse(mapThingToThingResponse(thing)), nil
}

func (s *ThingService) UpdateThing(ctx context.Context, request api.UpdateThingRequestObject) (api.UpdateThingResponseObject, error) {
	err := s.store.UpdateThing(helpers.Thing{
		UUID:  request.Uuid.String(),
		Score: request.Body.Score,
	})
	if err == helpers.ErrNotFound {
		return api.UpdateThing404JSONResponse{}, nil
	}
	if err != nil {
		return api.UpdateThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.UpdateThing204Response{}, nil
}

func (s *ThingService) ListThings(ctx context.Context, request api.ListThingsRequestObject) (api.ListThingsResponseObject, error) {
	limit := 10
	offset := 0
	if request.Params.Page != nil {
		offset = *request.Params.Page * limit
	}
	things, total := s.store.ListThings(offset, limit)
	response := api.ListThings200JSONResponse{Total: total, Things: []api.ThingResponse{}}
	for _, thing := range things {
		response.Things = append(response.Things, mapThingToThingResponse(thing))
	}
	return response, nil
}

func mapThingToThingResponse(thing helpers.Thing) api.ThingResponse {
	return api.ThingResponse{
		Created: thing.Created,
		Name:    thing.Name,
		Rank:    thing.Rank,
		Rating:  thing.Rating,
		Score:   thing.Score,
		Type:    api.ThingType(thing.Type),
		Uuid:    uuid.MustParse(thing.UUID),
	}
}
