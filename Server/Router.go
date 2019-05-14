package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(webHandler.notFound)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	//router.PathPrefix("/app/").Handler(http.StripPrefix("/app/", http.FileServer(http.Dir("./web/"))))
	//router.PathPrefix("/app").Handler(http.StripPrefix("/app", http.FileServer(http.Dir("./web/"))))

	return router
}
