// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeDeleteAPIV1NoteResponse(resp *http.Response) (res DeleteAPIV1NoteRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &DeleteAPIV1NoteOK{}, nil
	case 400:
		// Code 400.
		return &DeleteAPIV1NoteBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeDeleteAPIV1UserResponse(resp *http.Response) (res DeleteAPIV1UserRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &DeleteAPIV1UserOK{}, nil
	case 400:
		// Code 400.
		return &DeleteAPIV1UserBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetAPIV1NoteResponse(resp *http.Response) (res GetAPIV1NoteRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response Note
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &GetAPIV1NoteBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeGetAPIV1UserResponse(resp *http.Response) (res GetAPIV1UserRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response User
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &GetAPIV1UserBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeListAPIV1NoteResponse(resp *http.Response) (res ListAPIV1NoteRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response ListAPIV1NoteOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			// Validate response.
			if err := func() error {
				if err := response.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "validate")
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &ListAPIV1NoteBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodePostAPIV1NoteResponse(resp *http.Response) (res PostAPIV1NoteRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &PostAPIV1NoteOK{}, nil
	case 400:
		// Code 400.
		return &PostAPIV1NoteBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodePostAPIV1UserResponse(resp *http.Response) (res PostAPIV1UserRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		var wrapper PostAPIV1UserOK
		h := uri.NewHeaderDecoder(resp.Header)
		// Parse "Set-Cookie" header.
		{
			cfg := uri.HeaderParameterDecodingConfig{
				Name:    "Set-Cookie",
				Explode: false,
			}
			if err := func() error {
				if err := h.HasParam(cfg); err == nil {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotSetCookieVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							wrapperDotSetCookieVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.SetCookie.SetTo(wrapperDotSetCookieVal)
						return nil
					}); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return res, errors.Wrap(err, "parse Set-Cookie header")
			}
		}
		return &wrapper, nil
	case 400:
		// Code 400.
		return &PostAPIV1UserBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodePutAPIV1NoteResponse(resp *http.Response) (res PutAPIV1NoteRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &PutAPIV1NoteOK{}, nil
	case 400:
		// Code 400.
		return &PutAPIV1NoteBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodePutAPIV1UserResponse(resp *http.Response) (res PutAPIV1UserRes, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &PutAPIV1UserOK{}, nil
	case 400:
		// Code 400.
		return &PutAPIV1UserBadRequest{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
