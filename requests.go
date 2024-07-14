package main

import (
	"github.com/vanshcodes/go-netflix-backend/application/movies"
	"github.com/vanshcodes/go-netflix-backend/customtypes"
	// Types "github.com/vanshcodes/go-netflix-backend/types"
)

var requests = []customtypes.ClientRequest{

	{
		Path: "/movies",
		RequestDetails: []customtypes.RequestDetails{
			{
				RequestType:     "GET",
				RequestPassedTo: movies.GetAllMovies,
			},
		},
	},
}
