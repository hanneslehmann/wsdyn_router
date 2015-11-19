package main
//curl -X GET http://127.0.0.1:5552/movie/tt0082971
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie Struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

var movies = map[string]*Movie{
	"tt0076759": &Movie{Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
	"tt0082971": &Movie{Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	router.HandleFunc("/movie/{imdbKey}", handleMovie).Methods("GET")
	http.ListenAndServe(":5552", router)
}

func handleMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	imdbKey := vars["imdbKey"]

	log.Println("Request for:", imdbKey)

	if movie, ok := movies[imdbKey]; ok {
		outgoingJSON, error := json.Marshal(movie)

		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(res, string(outgoingJSON))
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, string("Movie not found"))
	}
}

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	outgoingJSON, error := json.Marshal(movies)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}
