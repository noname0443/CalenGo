// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetAPIV1NoteParams is parameters of get-api-v1-note operation.
type GetAPIV1NoteParams struct {
	Note string
}

func unpackGetAPIV1NoteParams(packed middleware.Parameters) (params GetAPIV1NoteParams) {
	{
		key := middleware.ParameterKey{
			Name: "note",
			In:   "path",
		}
		params.Note = packed[key].(string)
	}
	return params
}

func decodeGetAPIV1NoteParams(args [1]string, argsEscaped bool, r *http.Request) (params GetAPIV1NoteParams, _ error) {
	// Decode path: note.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "note",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Note = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "note",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetAPIV1UserParams is parameters of get-api-v1-user operation.
type GetAPIV1UserParams struct {
	User string
}

func unpackGetAPIV1UserParams(packed middleware.Parameters) (params GetAPIV1UserParams) {
	{
		key := middleware.ParameterKey{
			Name: "user",
			In:   "path",
		}
		params.User = packed[key].(string)
	}
	return params
}

func decodeGetAPIV1UserParams(args [1]string, argsEscaped bool, r *http.Request) (params GetAPIV1UserParams, _ error) {
	// Decode path: user.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "user",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.User = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "user",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}