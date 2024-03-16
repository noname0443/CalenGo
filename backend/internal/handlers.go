package internal

import (
	"context"

	api "github.com/noname0443/CalenGo/backend/api"
	ht "github.com/ogen-go/ogen/http"
	"github.com/sirupsen/logrus"
)

func (*App) GetAPIV1Note(ctx context.Context, params api.GetAPIV1NoteParams) (r api.GetAPIV1NoteRes, _ error) {
	logrus.Info("Request!")
	return r, ht.ErrNotImplemented
}
