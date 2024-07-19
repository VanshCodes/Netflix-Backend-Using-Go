package sqldatabasemanagers

import (
	"log"

	"github.com/vanshcodes/go-netflix-backend/db/models"
	"github.com/vanshcodes/go-netflix-backend/sqldb"
)

func GetAllMovies() models.Movies {

	query := "SELECT * FROM movies"
	var movies models.Movies

	err := sqldb.CLIENT_GLOBAL.Select(&movies, query)
	if err != nil {
		log.Fatal(err)
	}
	print(movies)
	return movies
}
