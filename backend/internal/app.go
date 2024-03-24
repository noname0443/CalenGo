package internal

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	api "github.com/noname0443/CalenGo/backend/api"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type App struct {
	api.UnimplementedHandler
	port string
	db   *sqlx.DB
}

func NewApp(port string, db *sqlx.DB) *App {
	return &App{
		port: port,
		db:   db,
	}
}

func (app *App) Run() {
	serv, err := api.NewServer(app, &SenHandler{}, []api.ServerOption{}...)
	if err != nil {
		logrus.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Authorization"},
	})

	logrus.Info("Server starting!")
	if err := http.ListenAndServe(app.port, c.Handler(serv)); err != nil {
		logrus.Fatal(err)
	}
}
