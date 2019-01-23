package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllMoviesEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func CreateMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndpoint).Methods("GET")
	r.HandleFunc("/movies", FindMovieEndpoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndpoint).Methods("POST")
	if err := http.ListenAndServe(":3003", r); err != nil {
		log.Fatal(err)
	}
}
