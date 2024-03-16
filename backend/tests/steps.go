package tests

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/noname0443/CalenGo/backend/api"
	"github.com/noname0443/CalenGo/backend/internal"
)

func NewFeatureManager() (*FeatureManager, error) {
	return &FeatureManager{}, nil
}

func (fm *FeatureManager) StepStartServerOn(ctx context.Context, ip string) (context.Context, error) {
	app := internal.NewApp(ip)
	fm.ip = ip
	go app.Run()
	return ctx, nil
}

func (fm *FeatureManager) StepIsWorkingOn(ctx context.Context) (context.Context, error) {
	err := checkServerIsWorking(fm.ip)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (fm *FeatureManager) StepGetNote(ctx context.Context, note string) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	resp, err := client.GetAPIV1Note(ctx, api.GetAPIV1NoteParams{
		Note: note,
	})
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepPostNote(ctx context.Context, data *godog.DocString) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	var note api.Note
	err = json.Unmarshal([]byte(data.Content), &note)
	if err != nil {
		return ctx, err
	}

	resp, err := client.PostAPIV1Note(ctx, api.NewOptNote(note))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepPutNote(ctx context.Context, data *godog.DocString) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	var note api.Note
	err = json.Unmarshal([]byte(data.Content), &note)
	if err != nil {
		return ctx, err
	}

	resp, err := client.PutAPIV1Note(ctx, api.NewOptNote(note))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepDeleteNote(ctx context.Context, data *godog.DocString) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	var note api.Note
	err = json.Unmarshal([]byte(data.Content), &note)
	if err != nil {
		return ctx, err
	}

	resp, err := client.DeleteAPIV1Note(ctx, api.NewOptNote(note))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepGetUser(ctx context.Context, user string) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	resp, err := client.GetAPIV1User(ctx, api.GetAPIV1UserParams{
		User: user,
	})
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepPostUser(ctx context.Context, data *godog.DocString) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	var user api.User
	err = json.Unmarshal([]byte(data.Content), &user)
	if err != nil {
		return ctx, err
	}

	resp, err := client.PostAPIV1User(ctx, api.NewOptUser(user))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepPutUser(ctx context.Context, data *godog.DocString) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	var user api.User
	err = json.Unmarshal([]byte(data.Content), &user)
	if err != nil {
		return ctx, err
	}

	resp, err := client.PutAPIV1User(ctx, api.NewOptUser(user))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepDeleteUser(ctx context.Context, data *godog.DocString) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	var user api.User
	err = json.Unmarshal([]byte(data.Content), &user)
	if err != nil {
		return ctx, err
	}

	resp, err := client.DeleteAPIV1User(ctx, api.NewOptUser(user))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) iGotLine(ctx context.Context, line string) (context.Context, error) {
	if line != fmt.Sprint(fm.lastResult) {
		return ctx, fmt.Errorf(fmt.Sprintf("%s not equal to %s", line, fm.lastResult))
	}
	return ctx, nil
}

type FeatureManager struct {
	lastResult interface{}
	ip         string
}
