package customtypes

import "net/http"

type RequestDetails struct {
	RequestType     string // GET, POST, PUT, DELETE
	RequestPassedTo func(http.ResponseWriter, *http.Request)
}

type ClientRequest struct {
	Path           string
	RequestDetails []RequestDetails
}
