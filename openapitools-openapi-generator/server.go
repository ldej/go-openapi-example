package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/ldej/go-openapi-example/helpers"
	openapi "github.com/ldej/go-openapi-example/openapitools-openapi-generator/gen/go"
)

func main() {

	// Create your usual storage/service/handler layers
	store := helpers.NewThingsDatastore()
	service := NewThingAPIService(store)
	controller := openapi.NewThingAPIController(service)
	router := openapi.NewRouter(controller)

	// Add swagger UI endpoints
	router.Get("/swagger/doc.yaml", func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.ReadFile("./gen/api/openapi.yaml")
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
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}

type ThingAPIService struct {
	store *helpers.ThingsDatastore
}

func NewThingAPIService(store *helpers.ThingsDatastore) openapi.ThingAPIServicer {
	return &ThingAPIService{store: store}
}

func (s *ThingAPIService) CreateThing(ctx context.Context, createThingRequest openapi.CreateThingRequest) (openapi.ImplResponse, error) {
	thing := helpers.Thing{
		UUID:    uuid.NewString(),
		Name:    createThingRequest.Name,
		Rank:    createThingRequest.Rank,
		Rating:  createThingRequest.Rating,
		Score:   createThingRequest.Score,
		Type:    string(createThingRequest.Type),
		Created: time.Now().UTC(),
	}
	err := s.store.StoreThing(thing)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	return openapi.Response(http.StatusOK, mapThingToThingResponse(thing)), nil
}

func (s *ThingAPIService) DeleteThing(ctx context.Context, uuid string) (openapi.ImplResponse, error) {
	err := s.store.DeleteThing(uuid)
	if err == helpers.ErrNotFound {
		return openapi.Response(http.StatusNotFound, nil), nil
	}
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	return openapi.Response(http.StatusNoContent, nil), nil
}

func (s *ThingAPIService) GetThing(ctx context.Context, uuid string) (openapi.ImplResponse, error) {
	thing, err := s.store.GetThing(uuid)
	if err == helpers.ErrNotFound {
		return openapi.Response(http.StatusNotFound, nil), nil
	}
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), err
	}
	return openapi.Response(http.StatusOK, mapThingToThingResponse(thing)), nil
}

func (s *ThingAPIService) ListThings(ctx context.Context, page int32, keyword string) (openapi.ImplResponse, error) {
	limit := 10
	offset := 0
	if page > 0 {
		offset = int(page) * limit
	}

	things, total := s.store.ListThings(offset, limit)
	response := openapi.ListThingsResponse{
		Total:  int32(total),
		Things: []openapi.ThingResponse{},
	}
	for _, thing := range things {
		response.Things = append(response.Things, mapThingToThingResponse(thing))
	}
	return openapi.Response(http.StatusOK, response), nil
}

func (s *ThingAPIService) UpdateThing(ctx context.Context, uuid string, updateThingRequest openapi.UpdateThingRequest) (openapi.ImplResponse, error) {
	err := s.store.UpdateThing(
		helpers.Thing{
			UUID:  uuid,
			Score: updateThingRequest.Score,
		},
	)
	if err == helpers.ErrNotFound {
		return openapi.Response(http.StatusNotFound, nil), nil
	}
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), nil
	}
	return openapi.Response(http.StatusNoContent, nil), nil
}

func mapThingToThingResponse(thing helpers.Thing) openapi.ThingResponse {
	return openapi.ThingResponse{
		Uuid:    thing.UUID,
		Type:    openapi.ThingType(thing.Type),
		Name:    thing.Name,
		Rank:    thing.Rank,
		Score:   thing.Score,
		Rating:  thing.Rating,
		Created: thing.Created,
	}
}
