// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleCreateThingRequest handles createThing operation.
//
// Create a thing.
//
// POST /things
func (s *Server) handleCreateThingRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createThing"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/things"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateThing",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "CreateThing",
			ID:   "createThing",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityApiKey(ctx, "CreateThing", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "ApiKey",
					Err:              err,
				}
				recordError("Security:ApiKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	request, close, err := s.decodeCreateThingRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *ThingResponse
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "CreateThing",
			OperationID:   "createThing",
			Body:          request,
			Params:        middleware.Parameters{},
			Raw:           r,
		}

		type (
			Request  = OptCreateThingRequest
			Params   = struct{}
			Response = *ThingResponse
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateThing(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateThing(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeCreateThingResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleDeleteThingRequest handles deleteThing operation.
//
// Delete a thing.
//
// DELETE /things/{uuid}
func (s *Server) handleDeleteThingRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteThing"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/things/{uuid}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteThing",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DeleteThing",
			ID:   "deleteThing",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityApiKey(ctx, "DeleteThing", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "ApiKey",
					Err:              err,
				}
				recordError("Security:ApiKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeDeleteThingParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response DeleteThingRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "DeleteThing",
			OperationID:   "deleteThing",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "uuid",
					In:   "path",
				}: params.UUID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteThingParams
			Response = DeleteThingRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteThingParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DeleteThing(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.DeleteThing(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDeleteThingResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleGetThingRequest handles getThing operation.
//
// Get a single thing.
//
// GET /things/{uuid}
func (s *Server) handleGetThingRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getThing"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/things/{uuid}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetThing",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "GetThing",
			ID:   "getThing",
		}
	)
	params, err := decodeGetThingParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetThingRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "GetThing",
			OperationID:   "getThing",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "uuid",
					In:   "path",
				}: params.UUID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetThingParams
			Response = GetThingRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetThingParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetThing(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetThing(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeGetThingResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleListThingsRequest handles listThings operation.
//
// List things.
//
// GET /things
func (s *Server) handleListThingsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listThings"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/things"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListThings",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ListThings",
			ID:   "listThings",
		}
	)
	params, err := decodeListThingsParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response ListThingsRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "ListThings",
			OperationID:   "listThings",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "page",
					In:   "query",
				}: params.Page,
				{
					Name: "keyword",
					In:   "query",
				}: params.Keyword,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ListThingsParams
			Response = ListThingsRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListThingsParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListThings(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListThings(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeListThingsResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleUpdateThingRequest handles updateThing operation.
//
// Update a thing.
//
// PUT /things/{uuid}
func (s *Server) handleUpdateThingRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateThing"),
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/things/{uuid}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateThing",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "UpdateThing",
			ID:   "updateThing",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityApiKey(ctx, "UpdateThing", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "ApiKey",
					Err:              err,
				}
				recordError("Security:ApiKey", err)
				s.cfg.ErrorHandler(ctx, w, r, err)
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			recordError("Security", err)
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
	}
	params, err := decodeUpdateThingParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeUpdateThingRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response UpdateThingRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "UpdateThing",
			OperationID:   "updateThing",
			Body:          request,
			Params: middleware.Parameters{
				{
					Name: "uuid",
					In:   "path",
				}: params.UUID,
			},
			Raw: r,
		}

		type (
			Request  = OptUpdateThingRequest
			Params   = UpdateThingParams
			Response = UpdateThingRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateThingParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateThing(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateThing(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		if errRes, ok := errors.Into[*ErrorStatusCode](err); ok {
			encodeErrorResponse(errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeUpdateThingResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}