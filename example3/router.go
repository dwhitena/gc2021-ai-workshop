package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is used to pass information about a particular route.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a slice of Route.
type Routes []Route

// Define our routes.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AnswerQuestion",
		"POST",
		"/answer",
		AnswerQuestion,
	},
}

// NewRouter forms a new mux router, see https://github.com/gorilla/mux.
func NewRouter() *mux.Router {

	// Create a basic router.
	router := mux.NewRouter().StrictSlash(false)
	router.SkipClean(true)

	// Assign the handlers to run when endpoints are called.
	for _, route := range routes {

		// Create a handler function.
		var handler http.Handler
		handler = route.HandlerFunc

		// Wrap all current routes in the logger decorator to log out requests.
		handler = Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	router.NotFoundHandler = router.NewRoute().HandlerFunc(NotFound).GetHandler()
	return router
}
