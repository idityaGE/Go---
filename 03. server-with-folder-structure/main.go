package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Println("Server running at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r)) // Bind to 0.0.0.0
}
