package sqldb

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	_ "github.com/lib/pq" // <------------ here

	"github.com/jmoiron/sqlx"
	"github.com/vanshcodes/go-netflix-backend/customtypes"
)

const DATABASE_TYPE = "postgres"

var CLIENT_GLOBAL *sqlx.DB

var CONFIGURATION_GLOBAL customtypes.DatabaseConfiguration

func StartConnection(configuration customtypes.DatabaseConfiguration) (*sqlx.DB, customtypes.DatabaseConfiguration) {

	slog.Info("Connecting to PostgreSQL...")

	// Connect to database
	port, err := strconv.Atoi(configuration.Port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to port: %v\n", err)
		os.Exit(2)
	}
	client, err := sqlx.Connect(DATABASE_TYPE, fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%d host=%s", configuration.Username, configuration.Password, configuration.DatabaseName, port, configuration.Host))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	CLIENT_GLOBAL = client
	CONFIGURATION_GLOBAL = configuration
	slog.Info("Connected to PostgreSQL")
	return client, configuration

}
