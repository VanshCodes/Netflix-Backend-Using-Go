package models

type Movies []Movie

type Movie struct {
	Adult            bool             `json:"adult"  bson:"adult"`
	BackdropPath     string           `json:"backdrop_path"  bson:"backdrop_path"`
	GenreIDS         []int64          `json:"genre_ids"  bson:"genre_ids"`
	ID               int64            `json:"id"  bson:"id"`
	OriginalLanguage OriginalLanguage `json:"original_language"  bson:"original_language"`
	OriginalTitle    string           `json:"original_title"  bson:"original_title"`
	Overview         string           `json:"overview"  bson:"overview"`
	Popularity       float64          `json:"popularity"  bson:"popularity"`
	PosterPath       string           `json:"poster_path"  bson:"poster_path"`
	ReleaseDate      string           `json:"release_date"  bson:"release_date"`
	Title            string           `json:"title"  bson:"title"`
	Video            bool             `json:"video"  bson:"video"`
	VoteAverage      float64          `json:"vote_average"  bson:"vote_average"`
	VoteCount        int64            `json:"vote_count"  bson:"vote_count"`
}

type OriginalLanguage string

const (
	En OriginalLanguage = "en"
	Ja OriginalLanguage = "ja"
	Ko OriginalLanguage = "ko"
	Pt OriginalLanguage = "pt"
)
