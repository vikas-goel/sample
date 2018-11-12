package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
}

var people []Person

func main() {
	people = append(people, Person{ID: "1", Firstname: "Vikas", Lastname: "Goel"})
	people = append(people, Person{ID: "2", Firstname: "Ankita", Lastname: "Goel"})
	router := mux.NewRouter()
	serve_people(router);

	log.Fatal(http.ListenAndServe(":8000", router))
}

func serve_people(r *mux.Router) {
	r.HandleFunc("/people", GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}

	http.NotFound(w, r)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, person := range people {
		if person.ID == params["id"] {
			return
		}
	}

	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, person := range people {
		if person.ID == params["id"] {
			json.NewDecoder(r.Body).Decode(&people[index])
			return
		}
	}

	http.NotFound(w, r)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			return
		}
	}

	http.NotFound(w, r)
}
