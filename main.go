package main

import (
	"fmt"
	"log"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
    fmt.Println("homePage()")
}

func handleRequests() {
    Router := mux.NewRouter().StrictSlash(true)
    Router.HandleFunc("/", homePage)
    Router.HandleFunc("/students", getStudents)
    Router.HandleFunc("/students/{id}", getOneStudent)

    log.Fatal(http.ListenAndServe(":8080", Router))
}

type Student struct {
    Id string `json:"Id"`
    Name string `json:"Name"`
    Age int `json:"Age"`
    FavouriteSubject string `json:"FavouriteSubject"`
}

var Students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getStudents()")
    json.NewEncoder(w).Encode(Students)
}

func getOneStudent(w http.ResponseWriter, r *http.Request) {
    fmt.Println("getOneStudent()")

    vars := mux.Vars(r)
    fmt.Println("mux.Vars():", vars)

    requestedId := vars["id"]
    for i := range Students {
        if Students[i].Id == requestedId {
            json.NewEncoder(w).Encode(Students[i])
        }
    }
}

func main() {
    Students = []Student {
        { Id: "0", Name: "Student A", Age: 11, FavouriteSubject: "Maths" },
        { Id: "1", Name: "Student B", Age: 14, FavouriteSubject: "Geography" },
        { Id: "2", Name: "Student C", Age: 8, FavouriteSubject: "English" },
    }
	handleRequests()
}
