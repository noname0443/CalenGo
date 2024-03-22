// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DeleteAPIV1Note implements delete-api-v1-note operation.
//
// Your POST endpoint.
//
// DELETE /api/v1/note
func (UnimplementedHandler) DeleteAPIV1Note(ctx context.Context, req OptNote, params DeleteAPIV1NoteParams) (r DeleteAPIV1NoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteAPIV1User implements delete-api-v1-user operation.
//
// Your POST endpoint.
//
// DELETE /api/v1/user
func (UnimplementedHandler) DeleteAPIV1User(ctx context.Context, req OptUser, params DeleteAPIV1UserParams) (r DeleteAPIV1UserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAPIV1Note implements get-api-v1-note operation.
//
// Your GET endpoint.
//
// GET /api/v1/note/{note}
func (UnimplementedHandler) GetAPIV1Note(ctx context.Context, params GetAPIV1NoteParams) (r GetAPIV1NoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAPIV1User implements get-api-v1-user operation.
//
// Your GET endpoint.
//
// GET /api/v1/user/{user}
func (UnimplementedHandler) GetAPIV1User(ctx context.Context, params GetAPIV1UserParams) (r GetAPIV1UserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListAPIV1Note implements list-api-v1-note operation.
//
// Your GET endpoint.
//
// GET /api/v1/note
func (UnimplementedHandler) ListAPIV1Note(ctx context.Context, req OptListAPIV1NoteReq, params ListAPIV1NoteParams) (r ListAPIV1NoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostAPIV1Note implements post-api-v1-note operation.
//
// Your POST endpoint.
//
// POST /api/v1/note
func (UnimplementedHandler) PostAPIV1Note(ctx context.Context, req OptNote, params PostAPIV1NoteParams) (r PostAPIV1NoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostAPIV1User implements post-api-v1-user operation.
//
// Your POST endpoint.
//
// POST /api/v1/user
func (UnimplementedHandler) PostAPIV1User(ctx context.Context, req OptUser) (r PostAPIV1UserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutAPIV1Note implements put-api-v1-note operation.
//
// Your POST endpoint.
//
// PUT /api/v1/note
func (UnimplementedHandler) PutAPIV1Note(ctx context.Context, req OptNote, params PutAPIV1NoteParams) (r PutAPIV1NoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutAPIV1User implements put-api-v1-user operation.
//
// Your POST endpoint.
//
// PUT /api/v1/user
func (UnimplementedHandler) PutAPIV1User(ctx context.Context, req OptUser, params PutAPIV1UserParams) (r PutAPIV1UserRes, _ error) {
	return r, ht.ErrNotImplemented
}
