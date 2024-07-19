package models

type Movies []Movie

type Movie struct {
	Adult            bool             `db:"adult"  bson:"adult"`
	BackdropPath     string           `db:"backdrop_path"  bson:"backdrop_path"`
	GenreIDS         []int64          `db:"genre_ids" json:"genre_ids"  bson:"genre_ids"`
	ID               int64            `db:"id"  bson:"id"`
	OriginalLanguage OriginalLanguage `db:"original_language"  bson:"original_language"`
	OriginalTitle    string           `db:"original_title"  bson:"original_title"`
	Overview         string           `db:"overview"  bson:"overview"`
	Popularity       float64          `db:"popularity"  bson:"popularity" json:"popularity"`
	PosterPath       string           `db:"poster_path"  bson:"poster_path"`
	ReleaseDate      string           `db:"release_date"  bson:"release_date" `
	Title            string           `db:"title"  bson:"title"`
	Video            bool             `db:"video"  bson:"video"`
	VoteAverage      float64          `db:"vote_average"  bson:"vote_average"`
	VoteCount        int64            `db:"vote_count"  bson:"vote_count"`
}

type OriginalLanguage string

const (
	En OriginalLanguage = "en"
	Ja OriginalLanguage = "ja"
	Ko OriginalLanguage = "ko"
	Pt OriginalLanguage = "pt"
)
