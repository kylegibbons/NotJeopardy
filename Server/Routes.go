package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		webHandler.Index,
	},
	// Infrastructure
	Route{
		"Web Socket",
		"GET",
		"/ws",
		func(w http.ResponseWriter, r *http.Request) {
			serveWs(hub, w, r)
		},
	},
	Route{
		"RAW",
		"GET",
		"/raw",
		serveRAW,
	},
	Route{
		"Client Application",
		"GET",
		"/app",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./web/index.html")

			//http.StripPrefix("/app", http.FileServer(http.Dir("./web/"))).ServeHTTP(w, r)*/
		},
	},
	Route{
		"Client Application",
		"GET",
		"/app/{rest:.*}",
		func(w http.ResponseWriter, r *http.Request) {
			rest := mux.Vars(r)["rest"]

			//log.Printf(vars["rest"])

			path := fmt.Sprintf("./web/%s", rest)

			if _, err := os.Stat(path); !os.IsNotExist(err) {
				http.ServeFile(w, r, path)
				return
			}

			// If the file has a . in the name, we assume they are looking for a specific file
			// and not a route in the SPA, so we return a 404
			if strings.Contains(rest, ".") {
				w.WriteHeader(http.StatusNotFound)

				fmt.Fprint(w, "404 Not Found!\r\r")
				fmt.Fprintf(w, "%v", r.RequestURI)

				log.Printf("File not found: %s", path)

				return
			}

			http.ServeFile(w, r, "./web/index.html")

			//http.StripPrefix("/app", http.FileServer(http.Dir("./web/"))).ServeHTTP(w, r)
		},
	},
}
