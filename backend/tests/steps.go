package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/cucumber/godog"
	docker "github.com/fsouza/go-dockerclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/noname0443/CalenGo/backend/api"
	"github.com/noname0443/CalenGo/backend/internal"
	"github.com/noname0443/CalenGo/backend/utility"
)

type FeatureManager struct {
	lastResult   interface{}
	Credentials  api.BasicAuth
	ip           string
	DockerClient *docker.Client
}

func (fm *FeatureManager) BasicAuth(ctx context.Context, operationName string) (api.BasicAuth, error) {
	return fm.Credentials, nil
}

func NewFeatureManager() (*FeatureManager, error) {
	dockerClient, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	return &FeatureManager{
		DockerClient: dockerClient,
	}, nil
}

func (fm *FeatureManager) StepCleanup(ctx context.Context) (context.Context, error) {
	err := fm.DockerClient.RemoveContainer(docker.RemoveContainerOptions{
		ID:    "test-mysql-integration-backend",
		Force: true,
	})
	return ctx, err
}

func (fm *FeatureManager) StepStartServerOn(ctx context.Context, ip string) (context.Context, error) {
	config := &docker.Config{
		Hostname: "test-mysql-integration-backend",
		Image:    "mysql",
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", "root"),
		},
		ExposedPorts: map[docker.Port]struct{}{
			docker.Port(fmt.Sprintf("%d/tcp", 15155)): struct{}{},
		},
	}

	hostConfig := &docker.HostConfig{
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port(fmt.Sprintf("%d/tcp", 3306)): []docker.PortBinding{
				{
					HostPort: fmt.Sprintf("%d", 15155),
				},
			},
		},
	}

	resp, err := fm.DockerClient.CreateContainer(docker.CreateContainerOptions{
		Name:       "test-mysql-integration-backend",
		Config:     config,
		HostConfig: hostConfig,
	})
	if err != nil {
		return ctx, err
	}

	if err := fm.DockerClient.StartContainer(resp.ID, &docker.HostConfig{}); err != nil {
		return ctx, err
	}

	var db *sqlx.DB
	err = utility.DoRetry(func() error {
		db, err = sqlx.Connect("mysql", "root:root@(localhost:15155)/sys")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return ctx, err
	}

	err = utility.InitSQL(db)
	if err != nil {
		return ctx, err
	}

	app := internal.NewApp(ip, db)
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

func (fm *FeatureManager) StepSetCredentials(ctx context.Context, credentials string) (context.Context, error) {
	strings := strings.Split(credentials, ":")
	if len(strings) != 2 {
		return ctx, errors.New("incorrect credentials")
	}
	fm.Credentials.Username = strings[0]
	fm.Credentials.Password = strings[1]
	return ctx, nil
}

func (fm *FeatureManager) StepGetNote(ctx context.Context, note string) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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

func (fm *FeatureManager) StepListNotes(ctx context.Context, start string, end string) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
	if err != nil {
		return ctx, err
	}

	resp, err := client.ListAPIV1Note(ctx, api.NewOptListAPIV1NoteReq(api.ListAPIV1NoteReq{
		StartTime: start,
		EndTime:   end,
	}))
	if err != nil {
		fm.lastResult = err
	} else {
		fm.lastResult = resp
	}

	return ctx, nil
}

func (fm *FeatureManager) StepGetUser(ctx context.Context, user string) (context.Context, error) {
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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
	client, err := api.NewClient(fmt.Sprintf("http://%s", fm.ip), fm, []api.ClientOption{}...)
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

func (fm *FeatureManager) iGotData(ctx context.Context, data *godog.DocString) (context.Context, error) {
	UnmarshaledStruct := reflect.New(reflect.TypeOf(fm.lastResult)).Interface()
	err := json.Unmarshal([]byte(data.Content), UnmarshaledStruct)
	valueAsLastResult := reflect.Indirect(reflect.ValueOf(UnmarshaledStruct))
	lastResult := reflect.ValueOf(fm.lastResult)
	if err != nil {
		return ctx, err
	}
	if !reflect.DeepEqual(lastResult.Interface(), valueAsLastResult.Interface()) {
		return ctx, errors.New(fmt.Sprintf("%s not equal to %s. Types: %s:%s", lastResult, valueAsLastResult, lastResult.Type(), valueAsLastResult.Type()))
	}
	return ctx, err
}

func (fm *FeatureManager) iGotNoError(ctx context.Context) (context.Context, error) {
	err, ok := fm.lastResult.(error)
	if !ok {
		return ctx, nil
	}
	return ctx, err
}
