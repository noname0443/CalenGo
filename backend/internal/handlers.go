package internal

import (
	"context"

	api "github.com/noname0443/CalenGo/backend/api"
	"github.com/sirupsen/logrus"
)

func (app *App) GetAPIV1Note(ctx context.Context, params api.GetAPIV1NoteParams) (api.GetAPIV1NoteRes, error) {
	note := api.Note{}
	err := app.db.Get(&note, readNoteByName, params.Note)
	logrus.Info("GetAPIV1Note", params, err)
	if err != nil {
		return &api.GetAPIV1NoteBadRequest{}, err
	}
	return &note, err
}

func (app *App) PostAPIV1Note(ctx context.Context, req api.OptNote) (r api.PostAPIV1NoteRes, err error) {
	_, err = app.db.NamedExec(createNoteByName, req.Value)
	logrus.Info("PostAPIV1Note", req, err)
	if err != nil {
		return &api.PostAPIV1NoteBadRequest{}, err
	}
	return &api.PostAPIV1NoteOK{}, err
}
