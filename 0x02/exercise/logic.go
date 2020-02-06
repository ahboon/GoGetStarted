package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {

}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}