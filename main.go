package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hewo, this is howme. you requested: %s\n", r.URL.Path)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	log.Println("Starting server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server failed to start", err)
	}
}