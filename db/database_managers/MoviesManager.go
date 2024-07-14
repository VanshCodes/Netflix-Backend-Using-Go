package databasemanagers

import (
	"context"
	"log"

	"github.com/vanshcodes/go-netflix-backend/constants"
	"github.com/vanshcodes/go-netflix-backend/db"
	"github.com/vanshcodes/go-netflix-backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllMovies(filters []interface{}) models.Movies {

	bsonInterface := []interface{}{
		filters
	}

	// movies := make([]customtypes.Movie, 0)
	var movies models.Movies
	database := db.CLIENT_GLOBAL.Database(db.CONFIGURATION_GLOBAL.DatabaseName)
	collection := database.Collection(constants.COLLECTION_NAME)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &movies) // TODO: Fix this
	if err != nil {
		log.Fatal(err)
		return models.Movies{}
	}
	return movies
}
