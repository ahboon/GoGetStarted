package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.URL.Query()["name"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    } 
    name,_ := r.URL.Query()["name"]
    if name[0] == "" {
        		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "name is missing."})
		return
    }
    if r.URL.Query()["age"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    } 
    age,_ := r.URL.Query()["age"]
    if age[0] == "" {
        		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorMessage{Error: "age is missing."})
		return
    }
    output := GetString(name[0],age[0])
    json.NewEncoder(w).Encode(output)
}

func ServeService() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/helloworld", HelloWorld).Methods("GET", "OPTIONS")
	return router
}