package internal

import (
	"net/http"

	api "github.com/noname0443/CalenGo/backend/api"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type App struct {
	api.UnimplementedHandler
	port string
}

func NewApp(port string) *App {
	return &App{
		port: port,
	}
}

func (app *App) Run() {
	serv, err := api.NewServer(app, []api.ServerOption{}...)
	if err != nil {
		logrus.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"User-Agent"},
	})

	logrus.Info("Server starting!")
	if err := http.ListenAndServe(app.port, c.Handler(serv)); err != nil {
		logrus.Fatal(err)
	}
}
