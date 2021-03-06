package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		hhf_index,
	},
	Route{
		"getCount",
		"GET",
		"/getCount",
		hhf_getCount,
	},
	Route{
		"getCount2",
		"GET",
		"/getCount2",
		hhf_getCount2,
	},
}
