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

// func HelloWorld(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     output := GetString()
//     json.NewEncoder(w).Encode(output)
// }

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}