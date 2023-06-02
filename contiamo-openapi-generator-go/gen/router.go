package api

// This file is auto-generated, don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// ThingsHandler handles the operations of the 'Things' handler group.
type ThingsHandler interface {
	// CreateThing Create a thing
	CreateThing(w http.ResponseWriter, r *http.Request)
	// DeleteThing Delete a thing
	DeleteThing(w http.ResponseWriter, r *http.Request)
	// GetThing Get a single thing
	GetThing(w http.ResponseWriter, r *http.Request)
	// ListThings List things
	ListThings(w http.ResponseWriter, r *http.Request)
	// UpdateThing Update a thing
	UpdateThing(w http.ResponseWriter, r *http.Request)
}

// NewRouter creates a new router for the spec and the given handlers.
// Things API
//
// The Things API creates, reads, updates, lists and deletes things!
//
// 1.0
func NewRouter(
	thingsHandler ThingsHandler,
) http.Handler {

	r := chi.NewRouter()

	// 'Things' group

	// '/things'
	r.Options("/things", optionsHandlerFunc(
		http.MethodGet,
		http.MethodPost,
	))
	r.Get("/things", thingsHandler.ListThings)
	r.Post("/things", thingsHandler.CreateThing)

	// '/things/{uuid}'
	r.Options("/things/{uuid}", optionsHandlerFunc(
		http.MethodDelete,
		http.MethodGet,
		http.MethodPut,
	))
	r.Delete("/things/{uuid}", thingsHandler.DeleteThing)
	r.Get("/things/{uuid}", thingsHandler.GetThing)
	r.Put("/things/{uuid}", thingsHandler.UpdateThing)

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
