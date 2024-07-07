package main

import (
	"database/sql"
	"fmt"
	"github.com/japhy-tech/backend-test/config"
	"net/http"
	"os"
	"time"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/database_actions"
	"github.com/japhy-tech/backend-test/internal"
)

func main() {
	logger := charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		Formatter:       charmLog.TextFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "🧑‍💻 backend-test",
		Level:           charmLog.DebugLevel,
	})

	err := database_actions.InitMigrator(config.MysqlDSN)
	if err != nil {
		logger.Fatal(err.Error())
	}

	msg, err := database_actions.RunMigrate("up", 0)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info(msg)
	}

	db, err := sql.Open("mysql", config.MysqlDSN)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	db.SetMaxIdleConns(0)

	err = db.Ping()
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	logger.Info("Database connected")

	// Load breeds data from CSV
	err = database_actions.LoadBreedsFromCSV(db, "breeds.csv")
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Breeds data loaded")

	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	app := internal.NewApp(db, r, logger)
	err = app.Start()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to start the application: %v", err))
	}

	// =============================== Starting Msg ===============================
	logger.Info(fmt.Sprintf("Service started and listen on port %s", config.ApiPort))
}
