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

// handleDeleteAPIV1NoteRequest handles delete-api-v1-note operation.
//
// Your POST endpoint.
//
// DELETE /api/v1/note
func (s *Server) handleDeleteAPIV1NoteRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("delete-api-v1-note"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/api/v1/note"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteAPIV1Note",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "DeleteAPIV1Note",
			ID:   "delete-api-v1-note",
		}
	)
	params, err := decodeDeleteAPIV1NoteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeDeleteAPIV1NoteRequest(r)
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

	var response DeleteAPIV1NoteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "DeleteAPIV1Note",
			OperationSummary: "Your POST endpoint",
			OperationID:      "delete-api-v1-note",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = OptNote
			Params   = DeleteAPIV1NoteParams
			Response = DeleteAPIV1NoteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteAPIV1NoteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DeleteAPIV1Note(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.DeleteAPIV1Note(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteAPIV1NoteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteAPIV1UserRequest handles delete-api-v1-user operation.
//
// Your POST endpoint.
//
// DELETE /api/v1/user
func (s *Server) handleDeleteAPIV1UserRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("delete-api-v1-user"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/api/v1/user"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteAPIV1User",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "DeleteAPIV1User",
			ID:   "delete-api-v1-user",
		}
	)
	params, err := decodeDeleteAPIV1UserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeDeleteAPIV1UserRequest(r)
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

	var response DeleteAPIV1UserRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "DeleteAPIV1User",
			OperationSummary: "Your POST endpoint",
			OperationID:      "delete-api-v1-user",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = OptUser
			Params   = DeleteAPIV1UserParams
			Response = DeleteAPIV1UserRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteAPIV1UserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DeleteAPIV1User(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.DeleteAPIV1User(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteAPIV1UserResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetAPIV1NoteRequest handles get-api-v1-note operation.
//
// Your GET endpoint.
//
// GET /api/v1/note/{note}
func (s *Server) handleGetAPIV1NoteRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("get-api-v1-note"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/note/{note}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetAPIV1Note",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "GetAPIV1Note",
			ID:   "get-api-v1-note",
		}
	)
	params, err := decodeGetAPIV1NoteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetAPIV1NoteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "GetAPIV1Note",
			OperationSummary: "Your GET endpoint",
			OperationID:      "get-api-v1-note",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "note",
					In:   "path",
				}: params.Note,
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetAPIV1NoteParams
			Response = GetAPIV1NoteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetAPIV1NoteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetAPIV1Note(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetAPIV1Note(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetAPIV1NoteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetAPIV1UserRequest handles get-api-v1-user operation.
//
// Your GET endpoint.
//
// GET /api/v1/user/{user}
func (s *Server) handleGetAPIV1UserRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("get-api-v1-user"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/user/{user}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetAPIV1User",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "GetAPIV1User",
			ID:   "get-api-v1-user",
		}
	)
	params, err := decodeGetAPIV1UserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetAPIV1UserRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "GetAPIV1User",
			OperationSummary: "Your GET endpoint",
			OperationID:      "get-api-v1-user",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "user",
					In:   "path",
				}: params.User,
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetAPIV1UserParams
			Response = GetAPIV1UserRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetAPIV1UserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetAPIV1User(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetAPIV1User(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetAPIV1UserResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleListAPIV1NoteRequest handles list-api-v1-note operation.
//
// Your GET endpoint.
//
// GET /api/v1/note
func (s *Server) handleListAPIV1NoteRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("list-api-v1-note"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/api/v1/note"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListAPIV1Note",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "ListAPIV1Note",
			ID:   "list-api-v1-note",
		}
	)
	params, err := decodeListAPIV1NoteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeListAPIV1NoteRequest(r)
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

	var response ListAPIV1NoteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "ListAPIV1Note",
			OperationSummary: "Your GET endpoint",
			OperationID:      "list-api-v1-note",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = OptListAPIV1NoteReq
			Params   = ListAPIV1NoteParams
			Response = ListAPIV1NoteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackListAPIV1NoteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ListAPIV1Note(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ListAPIV1Note(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeListAPIV1NoteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePostAPIV1NoteRequest handles post-api-v1-note operation.
//
// Your POST endpoint.
//
// POST /api/v1/note
func (s *Server) handlePostAPIV1NoteRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("post-api-v1-note"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/note"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PostAPIV1Note",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "PostAPIV1Note",
			ID:   "post-api-v1-note",
		}
	)
	params, err := decodePostAPIV1NoteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodePostAPIV1NoteRequest(r)
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

	var response PostAPIV1NoteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "PostAPIV1Note",
			OperationSummary: "Your POST endpoint",
			OperationID:      "post-api-v1-note",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = OptNote
			Params   = PostAPIV1NoteParams
			Response = PostAPIV1NoteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPostAPIV1NoteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.PostAPIV1Note(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.PostAPIV1Note(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePostAPIV1NoteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePostAPIV1UserRequest handles post-api-v1-user operation.
//
// Your POST endpoint.
//
// POST /api/v1/user
func (s *Server) handlePostAPIV1UserRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("post-api-v1-user"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/api/v1/user"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PostAPIV1User",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "PostAPIV1User",
			ID:   "post-api-v1-user",
		}
	)
	request, close, err := s.decodePostAPIV1UserRequest(r)
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

	var response PostAPIV1UserRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "PostAPIV1User",
			OperationSummary: "Your POST endpoint",
			OperationID:      "post-api-v1-user",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = OptUser
			Params   = struct{}
			Response = PostAPIV1UserRes
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
				response, err = s.h.PostAPIV1User(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.PostAPIV1User(ctx, request)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePostAPIV1UserResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePutAPIV1NoteRequest handles put-api-v1-note operation.
//
// Your POST endpoint.
//
// PUT /api/v1/note
func (s *Server) handlePutAPIV1NoteRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("put-api-v1-note"),
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/api/v1/note"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutAPIV1Note",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "PutAPIV1Note",
			ID:   "put-api-v1-note",
		}
	)
	params, err := decodePutAPIV1NoteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodePutAPIV1NoteRequest(r)
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

	var response PutAPIV1NoteRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "PutAPIV1Note",
			OperationSummary: "Your POST endpoint",
			OperationID:      "put-api-v1-note",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = OptNote
			Params   = PutAPIV1NoteParams
			Response = PutAPIV1NoteRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPutAPIV1NoteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.PutAPIV1Note(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.PutAPIV1Note(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutAPIV1NoteResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handlePutAPIV1UserRequest handles put-api-v1-user operation.
//
// Your POST endpoint.
//
// PUT /api/v1/user
func (s *Server) handlePutAPIV1UserRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("put-api-v1-user"),
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/api/v1/user"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutAPIV1User",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
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
			Name: "PutAPIV1User",
			ID:   "put-api-v1-user",
		}
	)
	params, err := decodePutAPIV1UserParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodePutAPIV1UserRequest(r)
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

	var response PutAPIV1UserRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "PutAPIV1User",
			OperationSummary: "Your POST endpoint",
			OperationID:      "put-api-v1-user",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "credentials",
					In:   "cookie",
				}: params.Credentials,
			},
			Raw: r,
		}

		type (
			Request  = OptUser
			Params   = PutAPIV1UserParams
			Response = PutAPIV1UserRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPutAPIV1UserParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.PutAPIV1User(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.PutAPIV1User(ctx, request, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutAPIV1UserResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
