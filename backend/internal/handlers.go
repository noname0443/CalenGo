package internal

import (
	"context"

	api "github.com/noname0443/CalenGo/backend/api"
	"github.com/sirupsen/logrus"
)

func (app *App) GetAPIV1Note(ctx context.Context, params api.GetAPIV1NoteParams) (api.GetAPIV1NoteRes, error) {
	note := api.Note{}
	err := app.db.Get(&note, readNoteByName, params.Note)
	logrus.Debug("GetAPIV1Note", params, err)
	if err != nil {
		return &api.GetAPIV1NoteBadRequest{}, err
	}
	return &note, err
}

func (app *App) PostAPIV1Note(ctx context.Context, req api.OptNote) (api.PostAPIV1NoteRes, error) {
	auth := ctx.Value("credentials").(api.BasicAuth)
	userId, _, err := app.GetUserByCredentials(auth)
	logrus.Debug("PostAPIV1Note", req, auth, err)
	if err != nil {
		return &api.PostAPIV1NoteBadRequest{}, err
	}

	_, err = app.db.Exec(createNoteByName, req.Value.Name, req.Value.Description, req.Value.StartTime, req.Value.EndTime, userId)
	if err != nil {
		return &api.PostAPIV1NoteBadRequest{}, err
	}
	return &api.PostAPIV1NoteOK{}, err
}

func (app *App) ListAPIV1Note(ctx context.Context, req api.OptListAPIV1NoteReq) (api.ListAPIV1NoteRes, error) {
	auth := ctx.Value("credentials").(api.BasicAuth)
	userId, _, err := app.GetUserByCredentials(auth)
	if err != nil {
		return &api.ListAPIV1NoteBadRequest{}, err
	}

	notes := []api.Note{}
	err = app.db.Select(&notes, readNotesByTime, req.Value.StartTime, req.Value.EndTime, userId)
	logrus.Debug("ListAPIV1Note", notes, req, err)
	if err != nil {
		return &api.ListAPIV1NoteBadRequest{}, err
	}

	resp := api.ListAPIV1NoteOKApplicationJSON(notes)
	return &resp, err
}

func (app *App) PostAPIV1User(ctx context.Context, req api.OptUser) (api.PostAPIV1UserRes, error) {
	_, err := app.db.NamedExec(createUser, req.Value)
	logrus.Debug("PostAPIV1User", req, err)
	if err != nil {
		return &api.PostAPIV1UserBadRequest{}, err
	}
	return &api.PostAPIV1UserOK{}, err
}

func (app *App) GetAPIV1User(ctx context.Context, params api.GetAPIV1UserParams) (api.GetAPIV1UserRes, error) {
	auth := ctx.Value("credentials").(api.BasicAuth)
	logrus.Debug("GetAPIV1User", params, auth)
	_, user, err := app.GetUserByCredentials(auth)
	return user, err
}

func (app *App) GetUserByCredentials(credentials api.BasicAuth) (int64, *api.User, error) {
	row := app.db.QueryRow(readUser, credentials.Username, credentials.Password)

	user := api.User{}
	var num int64
	err := row.Scan(&num, &user.Name, &user.Password)
	return num, &user, err
}
