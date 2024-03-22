package tests

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/sirupsen/logrus"
)

type logWriterType struct{ t *testing.T }

func (l logWriterType) Write(p []byte) (n int, err error) {
	l.t.Logf(string(p))
	return len(p), nil
}

func TestFeatures(t *testing.T) {
	logrus.SetOutput(logWriterType{t: t})
	logrus.SetLevel(logrus.DebugLevel)
	fm, err := NewFeatureManager()
	if err != nil {
		logrus.Fatal(err)
	}

	suite := godog.TestSuite{
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			InitializeScenario(fm, sc)
		},
		Options: &godog.Options{
			Format:        "pretty",
			Paths:         []string{"features"},
			StopOnFailure: true,
			TestingT:      t,
		},
	}
	t.Cleanup(func() {
		fm.StepCleanup(context.Background())
	})

	if suite.Run() != 0 {
		logrus.Info("non-zero status returned, failed to run feature tests")
		_, err := fm.StepCleanup(context.Background())
		logrus.Error(err)
	}
}

func InitializeScenario(fm *FeatureManager, sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})
	sc.Step(`^I start the server on "([^"]*)"$`, fm.StepStartServerOn)
	sc.Step(`^I set credentials "([^"]*)"$`, fm.StepSetCredentials)
	sc.Step(`^The server is working$`, fm.StepIsWorkingOn)

	sc.Step(`^I GET note /api/v1/note/([^"]*)$`, fm.StepGetNote)
	sc.Step(`^I POST note /api/v1/note$`, fm.StepPostNote)
	sc.Step(`^I PUT note /api/v1/note$`, fm.StepPutNote)
	sc.Step(`^I DELETE note /api/v1/note$`, fm.StepDeleteNote)
	sc.Step(`^I LIST notes >= ([^"]*) and <= ([^"]*)$`, fm.StepListNotes)

	sc.Step(`^I GET user /api/v1/user/([^"]*)$`, fm.StepGetUser)
	sc.Step(`^I POST user /api/v1/user$`, fm.StepPostUser)
	sc.Step(`^I PUT user /api/v1/user$`, fm.StepPutUser)
	sc.Step(`^I DELETE user /api/v1/user$`, fm.StepDeleteUser)

	sc.Step(`^I got line "([^"]*)"$`, fm.iGotLine)
	sc.Step(`^I got data$`, fm.iGotData)
	sc.Step(`^I got no error$`, fm.iGotNoError)
}
