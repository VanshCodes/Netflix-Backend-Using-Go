package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vanshcodes/go-netflix-backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Movie struct to represent a movie document
type Movie = models.Movie

func main2() {
	println("Hello")
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://192.168.1.107:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the collection
	collection = client.Database("Movies").Collection("movies")

	// Set up HTTP routes
	http.HandleFunc("/movies", getMoviesHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Handler function to retrieve all movies from MongoDB
func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	var movies []Movie

	// MongoDB find options (e.g., sort, filter) can be added here if needed
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error retrieving movies: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var movie Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Printf("Error decoding movie: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
