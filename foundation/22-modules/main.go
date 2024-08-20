package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go modules")

	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Serve the home page
	w.Write([]byte("<h1>Hello, World!</h1>"))
}
