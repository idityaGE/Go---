package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// some sample movies	(Get-Content .\main.go -Raw) -replace "`0", "" | Set-Content .\main.go
	movies = append(movies,
		Movie{ID: "1", ISBN: "123456", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}},
		Movie{ID: "2", ISBN: "654321", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}},
		Movie{ID: "3", ISBN: "987654", Title: "Movie Three", Director: &Director{Firstname: "Jane", Lastname: "Doe"}},
	)

	// Routes
	r.HandleFunc("/movies", getMovies).Methods("GET") // ✅
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") // ✅
	r.HandleFunc("/movies", createMovie).Methods("POST") // ✅
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT") // ✅
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // ✅

	fmt.Println("Server is running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/movies" {
	// 	http.Error(w, "400 Not Found", http.StatusBadRequest)
	// }
	// if r.Method != "GET" {
	// 	http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
	// 	return
	// }

	// setting header to json, so that the webbrowser or postman can understand that its an json
	w.Header().Set("Content-Type", "application/json")
	// encoding the movies slice into json to send json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // params --> /movies/{id}

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	// r.Body - return the body fron the request in json
	// json.NewDecoder used to decode the json same as JSON.parse in js
	// .Decode will put in movie struct
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Itoa --> int to str
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = params["id"]

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...) // delete exiting movie
		}
	}

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}