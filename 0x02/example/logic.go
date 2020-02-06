package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    output := GetString()
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}