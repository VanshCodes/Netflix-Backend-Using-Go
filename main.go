package main

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/vanshcodes/go-netflix-backend/constants"
	"github.com/vanshcodes/go-netflix-backend/customtypes"
	"github.com/vanshcodes/go-netflix-backend/db"
)

func SettingDatabase() {
	// Read the Database configuration from the file and set configuration
	var configuration customtypes.DatabaseConfiguration
	file, err := os.Open(constants.DATABASE_CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &configuration)
	db.StartConnection(configuration)
}

func main() {
	// slog.Info()
	slog.Info("Starting server at " + constants.PORT + " ...")
	router := mux.NewRouter()
	for _, path := range requests {
		for _, path_request := range path.RequestDetails {
			slog.Info("Adding handler for " + path_request.RequestType + " for " + path.Path)
			router.HandleFunc(path.Path, path_request.RequestPassedTo).Methods(path_request.RequestType)
		}
	}
	slog.Info("Setting database connection...")
	SettingDatabase()

	// slog.Info(constants.PORT)
	log.Fatal(http.ListenAndServe(constants.PORT, router))
}
