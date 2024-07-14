package movies

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"slices"
	"strconv"

	"github.com/vanshcodes/go-netflix-backend/db"
	databasemanagers "github.com/vanshcodes/go-netflix-backend/db/database_managers"
	"github.com/vanshcodes/go-netflix-backend/db/models"
	muxhelpers "github.com/vanshcodes/go-netflix-backend/mux_helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

func MovieQueryFilter(w http.ResponseWriter, r *http.Request, q func(w http.ResponseWriter, r *http.Request) *mongo.Cursor) {
	var movies models.Movies
	defer r.Body.Close()
	slog.Info(db.CONFIGURATION_GLOBAL.Host)
	cursor := q(w, r)
	err := cursor.All(context.TODO(), &movies)
	if err != nil {
		log.Fatal(err)
	}
	muxhelpers.JsonResponse(w, movies)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	slog.Info("HOST IS")
	movies := databasemanagers.GetAllMovies()
	movies = CheckAndUseFilters(movies, r)
	muxhelpers.JsonResponse(w, movies)
}

//	func GetMovieByGenre(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		genrestrID := r.URL.Query().Get("genre_id")
//		genreID, err := strconv.Atoi(genrestrID)
//		if err != nil {
//			log.Fatal(err)
//		}
//		movies := databasemanagers.GetMovieByGenre(int64(genreID))
//		muxhelpers.JsonResponse(w, movies)
//	}
func FilterMovieResultsByGenre(movieResult models.Movies, genreID int64) models.Movies {
	var movies models.Movies
	for _, movie := range movieResult {
		println(movie.GenreIDS[len(movie.GenreIDS)-1])
		if slices.Contains(movie.GenreIDS, genreID) {
			// println(slices.Contains(movie.GenreIDS, genreID))

			movies = append(movies, movie)
		}
	}
	return movies
}

func CheckAndUseFilters(movies models.Movies, r *http.Request) models.Movies {
	if r.URL.Query().Get("genre_id") != "" {

		genreID, err := strconv.Atoi(r.URL.Query().Get("genre_id"))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Genre: ", genreID)

		movies = FilterMovieResultsByGenre(movies, int64(genreID))
	}
	return movies
}
