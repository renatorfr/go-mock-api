package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type endpoint struct {
	Description string
	Pattern     string
	Methods     string
	Request     request
	Response    response
}

type request struct {
	Verb string
	Body string
}

type response struct {
	StatusCode string
	Body       string
}

var endpoints map[string]endpoint

func main() {
	startServer()
}

func startServer() {
	router := mux.NewRouter()
	loadEndpoints()
	loadRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func buildRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func loadRoutes(router *mux.Router) {
	for path, endpoint := range endpoints {
		router.HandleFunc(path, defaultHandler).Methods(endpoint.Methods)
	}
}

func loadEndpoints() {
	endpoints = map[string]endpoint{
		"/": endpoint{
			"Default root",
			"",
			"GET",
			request{},
			response{Body: "default"},
		},
		"/teste1": endpoint{
			"Description 1",
			"",
			"GET",
			request{},
			response{Body: "bodiiii"},
		},
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(endpoints[r.URL.Path].Response.Body)
}
