// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateThing implements createThing operation.
	//
	// Create a thing.
	//
	// POST /things
	CreateThing(ctx context.Context, req OptCreateThingRequest) (*ThingResponse, error)
	// DeleteThing implements deleteThing operation.
	//
	// Delete a thing.
	//
	// DELETE /things/{uuid}
	DeleteThing(ctx context.Context, params DeleteThingParams) (DeleteThingRes, error)
	// GetThing implements getThing operation.
	//
	// Get a single thing.
	//
	// GET /things/{uuid}
	GetThing(ctx context.Context, params GetThingParams) (GetThingRes, error)
	// ListThings implements listThings operation.
	//
	// List things.
	//
	// GET /things
	ListThings(ctx context.Context, params ListThingsParams) (ListThingsRes, error)
	// UpdateThing implements updateThing operation.
	//
	// Update a thing.
	//
	// PUT /things/{uuid}
	UpdateThing(ctx context.Context, req OptUpdateThingRequest, params UpdateThingParams) (UpdateThingRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
