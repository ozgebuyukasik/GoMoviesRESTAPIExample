package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var moviesList []Movie
var directorsList []Director

func getMovies(responseWriter http.ResponseWriter, req *http.Request){
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(moviesList)
}

func getMovie(responseWriter http.ResponseWriter, req *http.Request){
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, movie := range moviesList {
		if movie.ID == params["id"]{
			json.NewEncoder(responseWriter).Encode(movie)
			return
		}
	}
}

func createMovie(responseWriter http.ResponseWriter, req *http.Request){
	responseWriter.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(req.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(10000))
	moviesList = append(moviesList, newMovie)
	json.NewEncoder(responseWriter).Encode(moviesList)
}

func updateMovie(responseWriter http.ResponseWriter, req *http.Request){
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, movie := range moviesList {
		if movie.ID == params["id"] {
			moviesList = append(moviesList[:index], moviesList[index+1:]... )
			var updatedMovie Movie
			_ = json.NewDecoder(req.Body).Decode(&updatedMovie)
			updatedMovie.ID = strconv.Itoa(rand.Intn(10000))
			moviesList = append(moviesList, updatedMovie)
			json.NewEncoder(responseWriter).Encode(moviesList)
			return
		}
	}
}

func deleteMovie(responseWriter http.ResponseWriter, req *http.Request){
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, movie := range moviesList{

		if movie.ID == params["id"]{
			moviesList = append(moviesList[:index], moviesList[index+1:]...)
			break
		}
	}
	json.NewEncoder(responseWriter).Encode(moviesList)
}


func main(){
	directorsList = append(directorsList, Director{Firstname: "Frank", Lastname: "Darabont"})
	directorsList = append(directorsList, Director{Firstname: "Francis", Lastname: "Coppola"})
	directorsList = append(directorsList, Director{Firstname: "Steven", Lastname: "Spielberg"})

	moviesList = append(moviesList, Movie{ID: "1", Isbn: "74085", Title: "The Shawshank Redemption", Director: &directorsList[0]})
	moviesList = append(moviesList, Movie{ID: "2", Isbn: "30854", Title: "The Godfather", Director: &directorsList[1]})
	moviesList = append(moviesList, Movie{ID: "3", Isbn: "84569", Title: "Schindler's List", Director: &directorsList[2]})

	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is started at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}