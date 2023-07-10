package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	routes := httprouter.New()

	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST"
	// respectively.
	routes.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	routes.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	routes.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return routes
}
