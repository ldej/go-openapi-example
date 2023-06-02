package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	api "github.com/ldej/go-openapi-example/contiamo-openapi-generator-go/gen"
	"github.com/ldej/go-openapi-example/helpers"
)

func main() {
	store := helpers.NewThingsDatastore()
	thingsHandler := NewThingsHandler(store)
	router := chi.NewRouter()

	router.Route("/api/v1/", func(r chi.Router) {
		r.Mount("/", api.NewRouter(thingsHandler))
	})

	// Add swagger UI endpoints
	router.Get("/swagger/doc.yaml", func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.ReadFile("./openapi3-modified.yaml")
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

func NewThingsHandler(store *helpers.ThingsDatastore) api.ThingsHandler {
	return &ThingsHandler{
		store: store,
	}
}

type ThingsHandler struct {
	store *helpers.ThingsDatastore
}

func (h *ThingsHandler) CreateThing(w http.ResponseWriter, r *http.Request) {
	var request api.CreateThingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if err := request.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	thing := helpers.Thing{
		UUID:    uuid.NewString(),
		Name:    request.GetName(),
		Rank:    request.GetRank(),
		Rating:  request.GetRating(),
		Score:   request.Score,
		Type:    string(request.Type),
		Created: time.Now().UTC(),
	}
	err := h.store.StoreThing(thing)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, http.StatusOK, mapThingToThingResponse(thing))
}

func (h *ThingsHandler) DeleteThing(w http.ResponseWriter, r *http.Request) {
	queryParameters := api.DeleteThingQueryParameters{
		Uuid: chi.URLParam(r, "uuid"),
	}
	if err := queryParameters.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	err := h.store.DeleteThing(queryParameters.GetUuid())
	if err == helpers.ErrNotFound {
		writeError(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, http.StatusNoContent, nil)
}

func (h *ThingsHandler) GetThing(w http.ResponseWriter, r *http.Request) {
	queryParameters := api.GetThingQueryParameters{
		Uuid: chi.URLParam(r, "uuid"),
	}
	if err := queryParameters.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	thing, err := h.store.GetThing(queryParameters.GetUuid())
	if err == helpers.ErrNotFound {
		writeError(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, http.StatusOK, mapThingToThingResponse(thing))
}

func (h *ThingsHandler) ListThings(w http.ResponseWriter, r *http.Request) {
	queryParameters := api.ListThingsQueryParameters{
		Page:    int32(stringToIntDefault(chi.URLParam(r, "page"), 1)),
		Keyword: chi.URLParam(r, "keyword"),
	}
	if err := queryParameters.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	limit := 10
	offset := int(queryParameters.Page) * limit
	things, total := h.store.ListThings(offset, limit)

	listThingsResponse := api.ListThingsResponse{
		Things: []api.ThingResponse{},
		Total:  int32(total),
	}
	for _, thing := range things {
		listThingsResponse.Things = append(listThingsResponse.Things, mapThingToThingResponse(thing))
	}
	writeResponse(w, http.StatusOK, listThingsResponse)
}

func (h *ThingsHandler) UpdateThing(w http.ResponseWriter, r *http.Request) {
	var request api.UpdateThingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if err := request.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	queryParameters := api.UpdateThingQueryParameters{
		Uuid: chi.URLParam(r, "uuid"),
	}
	if err := queryParameters.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	thing := helpers.Thing{
		UUID:  queryParameters.Uuid,
		Score: request.Score,
	}
	err := h.store.UpdateThing(thing)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeResponse(w, http.StatusNoContent, nil)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	e := api.Error{
		Code:    strconv.Itoa(statusCode),
		Message: err.Error(),
	}
	writeResponse(w, statusCode, e)
}

func writeResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}
}

func mapThingToThingResponse(thing helpers.Thing) api.ThingResponse {
	return api.ThingResponse{
		Created: thing.Created,
		Name:    thing.Name,
		Rank:    thing.Rank,
		Rating:  thing.Rating,
		Score:   thing.Score,
		Type:    api.ThingType(thing.Type),
		Uuid:    thing.UUID,
	}
}

func stringToIntDefault(s string, i int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return i
	}
	return v
}
