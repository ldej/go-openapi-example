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
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(swagger)
	})
	r.Handle("/swagger/", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		Path:    "/swagger/",
		SpecURL: "/swagger/doc.json",
	}, nil))

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

	// Create your usual storage/service/handler layers
	thingsStore := helpers.NewThingsDatastore()
	thingsHandler := NewThingsHandler(thingsStore)
	s := NewServer(thingsHandler)

	apiServer := api.HandlerWithOptions(
		api.NewStrictHandler(s, nil),
		api.ChiServerOptions{
			BaseURL:     "/api/v1",
			BaseRouter:  r,
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

func NewServer(thingsHandler *ThingsHandler) api.StrictServerInterface {
	return &server{
		ThingsHandler: thingsHandler,
	}
}

type server struct {
	*ThingsHandler
}

type ThingsHandler struct {
	store *helpers.ThingsDatastore
}

func NewThingsHandler(store *helpers.ThingsDatastore) *ThingsHandler {
	return &ThingsHandler{
		store: store,
	}
}

func (h *ThingsHandler) CreateThing(ctx context.Context, request api.CreateThingRequestObject) (api.CreateThingResponseObject, error) {
	thing := helpers.Thing{
		UUID:    uuid.NewString(),
		Name:    request.Body.Name,
		Rank:    request.Body.Rank,
		Rating:  helpers.ToFloat32(request.Body.Rating),
		Score:   helpers.ToFloat64(request.Body.Score),
		Type:    string(request.Body.Type),
		Created: time.Now().UTC(),
	}
	err := h.store.StoreThing(thing)
	if err != nil {
		return api.CreateThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.CreateThing200JSONResponse(mapThingToThingResponse(thing)), nil
}

func (h *ThingsHandler) DeleteThing(ctx context.Context, request api.DeleteThingRequestObject) (api.DeleteThingResponseObject, error) {
	err := h.store.DeleteThing(request.Uuid.String())
	if err == helpers.ErrNotFound {
		return api.DeleteThing404JSONResponse{}, nil
	}
	if err != nil {
		return api.DeleteThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.DeleteThing204Response{}, nil
}

func (h *ThingsHandler) GetThing(ctx context.Context, request api.GetThingRequestObject) (api.GetThingResponseObject, error) {
	thing, err := h.store.GetThing(request.Uuid.String())
	if err == helpers.ErrNotFound {
		return api.GetThing404JSONResponse{}, nil
	}
	if err != nil {
		return api.GetThingdefaultJSONResponse{Body: api.Error{Message: err.Error()}, StatusCode: http.StatusInternalServerError}, nil
	}
	return api.GetThing200JSONResponse(mapThingToThingResponse(thing)), nil
}

func (h *ThingsHandler) UpdateThing(ctx context.Context, request api.UpdateThingRequestObject) (api.UpdateThingResponseObject, error) {
	err := h.store.UpdateThing(helpers.Thing{
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

func (h *ThingsHandler) ListThings(ctx context.Context, request api.ListThingsRequestObject) (api.ListThingsResponseObject, error) {
	limit := 10
	offset := 0
	if request.Params.Page != nil {
		offset = *request.Params.Page * limit
	}
	things, total := h.store.ListThings(offset, limit)
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
