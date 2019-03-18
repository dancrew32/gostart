package main

import (
  "fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthCheck struct {
	Version string `json:"version"`
	Alive   bool   `json:"alive"`
}

type Foo struct {
	Bar string `json:"bar"`
	Wat bool   `json:"wat"`
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
  var foo Foo
  json.NewDecoder(r.Body).Decode(&foo)
  fmt.Fprintf(w, "%s and %t", foo.Bar, foo.Wat)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := HealthCheck{
		Alive:   true,
		Version: "abcdef",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/decode", DecodeHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
