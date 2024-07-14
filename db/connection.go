package db

import (
	"context"
	"log"
	"log/slog"

	"github.com/vanshcodes/go-netflix-backend/customtypes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE_TYPE = "mongodb"

var CLIENT_GLOBAL *mongo.Client

var CONFIGURATION_GLOBAL customtypes.DatabaseConfiguration

func StartConnection(configuration customtypes.DatabaseConfiguration) (*mongo.Client, customtypes.DatabaseConfiguration) {

	slog.Info("Connecting to MongoDB...")
	URI := DATABASE_TYPE + "://" + configuration.Host + ":" + configuration.Port + "/" + configuration.DatabaseName
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))

	if err != nil {
		log.Fatal("Error connecting to MongoDB")
		panic(err)
	}
	CLIENT_GLOBAL = client
	CONFIGURATION_GLOBAL = configuration
	slog.Info("Connected to MongoDB")

	return client, configuration
}
