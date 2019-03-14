package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homeController(w http.ResponseWriter, req *http.Request) {

}

func restaurantsController(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Incoming %v request to Restaurants", req.Method)

	params := mux.Vars(req)
	fmt.Printf("My params are %v", params)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", homeController)
	r.HandleFunc("/restaurants", restaurantsController).Methods("GET", "POST")
	http.Handle("/", r)

	http.ListenAndServe(":"+port, nil)

}
