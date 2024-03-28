package internal

import (
	"context"
	"errors"

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
	if req.Value.Name == "" {
		return &api.PostAPIV1NoteBadRequest{}, errors.New("name is empty")
	}
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

func (app *App) ConflictAPIV1(ctx context.Context) (api.ConflictAPIV1Res, error) {
	auth := ctx.Value("credentials").(api.BasicAuth)
	logrus.Debug("ConflictAPIV1", auth)
	id, _, err := app.GetUserByCredentials(auth)
	if err != nil {
		logrus.Debug("ConflictAPIV1", err)
		return &api.ConflictAPIV1BadRequest{}, err
	}

	rows, err := app.db.Query(readConflictedNotes, id)
	if err != nil {
		logrus.Debug("ConflictAPIV1", err)
		return &api.ConflictAPIV1BadRequest{}, err
	}

	items := make([]api.ConflictAPIV1OKItem, 0)
	for rows.Next() {
		first, second := api.Note{}, api.Note{}
		err = rows.Scan(&first.Name, &first.Description, &first.StartTime, &first.EndTime, &second.Name, &second.Description, &second.StartTime, &second.EndTime)
		if err != nil {
			logrus.Debug("ConflictAPIV1", err)
			return &api.ConflictAPIV1BadRequest{}, err
		}
		items = append(items, api.ConflictAPIV1OKItem{
			First:  first,
			Second: second,
		})
	}

	logrus.Debug("ConflictAPIV1", items, err)
	resp := api.ConflictAPIV1OKApplicationJSON(items)
	return &resp, err
}
