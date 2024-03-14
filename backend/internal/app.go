package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type App struct {
	fiberApp *fiber.App
	port     string
}

func NewApp(port string) *App {
	app := &App{
		fiberApp: fiber.New(),
		port:     port,
	}

	app.fiberApp.Get("/", func(c *fiber.Ctx) error {
		logrus.Info("Hello, world!")
		return c.SendString("Hello, World!")
	})

	return app
}

func (app *App) Run() {
	err := app.fiberApp.Listen(app.port)
	if err != nil {
		logrus.Fatal(err)
	}
}
