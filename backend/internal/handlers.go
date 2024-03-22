package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (app *App) PostAPIV1Note(ctx context.Context, req api.OptNote, data api.PostAPIV1NoteParams) (api.PostAPIV1NoteRes, error) {
	userId, _, err := app.GetUserByCredentials(data.Credentials)
	logrus.Debug("PostAPIV1Note", req, data, err)
	if err != nil {
		return &api.PostAPIV1NoteBadRequest{}, err
	}

	_, err = app.db.Exec(createNoteByName, req.Value.Name, req.Value.Description, req.Value.StartTime, req.Value.EndTime, userId)
	if err != nil {
		return &api.PostAPIV1NoteBadRequest{}, err
	}
	return &api.PostAPIV1NoteOK{}, err
}

func (app *App) ListAPIV1Note(ctx context.Context, req api.OptListAPIV1NoteReq, data api.ListAPIV1NoteParams) (api.ListAPIV1NoteRes, error) {
	notes := []api.Note{}
	err := app.db.Select(&notes, readNotesByTime, req.Value.StartTime, req.Value.EndTime)
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
	return &api.PostAPIV1UserOK{
		SetCookie: api.NewOptString(fmt.Sprintf("%s:%s", req.Value.Name, req.Value.Password)),
	}, err
}

func (app *App) GetAPIV1User(ctx context.Context, params api.GetAPIV1UserParams) (api.GetAPIV1UserRes, error) {
	logrus.Debug("GetAPIV1User", params)
	strings := strings.Split(params.Credentials, ":")
	if len(strings) != 2 {
		return nil, errors.New("incorrect credentials")
	}

	if params.User != strings[0] {
		return nil, errors.New("incorrect cookie's user or the user in request")
	}
	_, user, err := app.GetUserByCredentials(params.Credentials)
	return user, err
}

func (app *App) GetUserByCredentials(credentials string) (int64, *api.User, error) {
	strings := strings.Split(credentials, ":")
	if len(strings) != 2 {
		return 0, nil, errors.New("incorrect credentials")
	}

	username, password := strings[0], strings[1]
	row := app.db.QueryRow(readUser, username, password)

	user := api.User{}
	var num int64
	err := row.Scan(&num, &user.Name, &user.Password)
	return num, &user, err
}
