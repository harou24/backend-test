package internal

import (
	"database/sql"
	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/config"
	"github.com/japhy-tech/backend-test/internal/api"
	"github.com/japhy-tech/backend-test/internal/infrastructure/persistence"
	"net"
	"net/http"
)

type App struct {
	logger *charmLog.Logger
	db     *sql.DB
	router *mux.Router
}

func NewApp(db *sql.DB, router *mux.Router, logger *charmLog.Logger) *App {
	return &App{
		logger: logger,
		db:     db,
		router: router,
	}
}

func (a *App) Start() error {
	breedRepository := persistence.NewMysqlBreedRepository(a.db)
	breedHandler := api.NewBreedHandler(breedRepository)
	a.router.Use(api.EnableCORS)

	api.RegisterRoutes(a.router.PathPrefix("/v1").Subrouter(), breedHandler)

	err := http.ListenAndServe(
		net.JoinHostPort("", config.ApiPort),
		a.router,
	)
	if err != nil {
		a.logger.Error("Failed to start server: ", err)
		return err
	}
	return nil
}
