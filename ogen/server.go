package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/ldej/go-openapi-example/helpers"
	api "github.com/ldej/go-openapi-example/ogen/gen"
)

func main() {
	store := helpers.NewThingsDatastore()
	service := NewThingService(store)
	securityHandler := NewSecurityHandler()

	apiPathPrefix := "/api/v1"

	server, err := api.NewServer(service, securityHandler, api.WithPathPrefix(apiPathPrefix))
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Mount(apiPathPrefix, server)

	// Add swagger UI endpoints
	router.Get("/swagger/doc.yaml", func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.ReadFile("../openapi3.yaml")
		w.Write(file)
	})
	router.Handle("/swagger/", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		Path:    "/swagger/",
		SpecURL: "/swagger/doc.yaml",
	}, nil))

	addr := ":8000"
	httpServer := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("Server listening on", addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}

type ThingService struct {
	store *helpers.ThingsDatastore
}

func NewThingService(store *helpers.ThingsDatastore) api.Handler {
	return &ThingService{store: store}
}

func (s *ThingService) CreateThing(ctx context.Context, req api.OptCreateThingRequest) (*api.ThingResponse, error) {
	thing := helpers.Thing{
		UUID:    uuid.NewString(),
		Name:    req.Value.Name,
		Rank:    req.Value.Rank,
		Rating:  req.Value.Rating.Or(0),
		Score:   req.Value.Score.Or(0),
		Type:    string(req.Value.Type),
		Created: time.Now().UTC(),
	}
	if err := s.store.StoreThing(thing); err != nil {
		return nil, err
	}
	return mapThingToThingResponse(thing), nil
}

func (s *ThingService) DeleteThing(ctx context.Context, params api.DeleteThingParams) (api.DeleteThingRes, error) {
	err := s.store.DeleteThing(params.UUID.String())
	if err == helpers.ErrNotFound {
		return &api.DeleteThingNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}
	return &api.DeleteThingNoContent{}, nil
}

func (s *ThingService) GetThing(ctx context.Context, params api.GetThingParams) (api.GetThingRes, error) {
	thing, err := s.store.GetThing(params.UUID.String())
	if err == helpers.ErrNotFound {
		return &api.GetThingNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}
	return mapThingToThingResponse(thing), nil
}

func (s *ThingService) ListThings(ctx context.Context, params api.ListThingsParams) (api.ListThingsRes, error) {
	page := 0
	limit := 10
	value, set := params.Page.Get()
	if set {
		page = value
	}
	things, total := s.store.ListThings(page-1*limit, limit)

	response := api.ListThingsResponse{
		Things: []api.ThingResponse{},
		Total:  total,
	}
	for _, thing := range things {
		response.Things = append(response.Things, *mapThingToThingResponse(thing))
	}
	return &response, nil
}

func (s *ThingService) UpdateThing(ctx context.Context, req api.OptUpdateThingRequest, params api.UpdateThingParams) (api.UpdateThingRes, error) {
	err := s.store.UpdateThing(helpers.Thing{
		UUID:  params.UUID.String(),
		Score: req.Value.Score,
	})
	if err == helpers.ErrNotFound {
		return &api.UpdateThingNotFound{}, nil
	}
	if err != nil {
		return nil, err
	}
	return &api.UpdateThingNoContent{}, nil
}

func (s *ThingService) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Code:    strconv.Itoa(http.StatusInternalServerError),
			Message: err.Error(),
		},
	}
}

type SecurityHandler struct{}

func NewSecurityHandler() api.SecurityHandler {
	return &SecurityHandler{}
}

func (s SecurityHandler) HandleApiKey(ctx context.Context, operationName string, t api.ApiKey) (context.Context, error) {
	log.Println("API key", t.GetAPIKey(), "is trying to access", operationName)
	return ctx, nil
}

func mapThingToThingResponse(thing helpers.Thing) *api.ThingResponse {
	return &api.ThingResponse{
		UUID:    uuid.MustParse(thing.UUID),
		Type:    api.ThingType(thing.Type),
		Name:    thing.Name,
		Rank:    thing.Rank,
		Score:   thing.Score,
		Rating:  thing.Rating,
		Created: thing.Created,
	}
}
