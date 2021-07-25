package main

import (
	"fmt"
	"log"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!")
    fmt.Println("homePage()")
}

func handleRequests() {
    Router := mux.NewRouter().StrictSlash(true)
    Router.HandleFunc("/", homePage)
    Router.HandleFunc("/list-students", getStudents)
    Router.HandleFunc("/create-student", createStudent).Methods("POST")
    Router.HandleFunc("/view-student/{id}", getOneStudent)

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

// via POST
// curl -X POST -H 'Content-Type: application/json' -d '{"Id": "3", "Name": "Student D", "Age": 12, "FavouriteSubject": "History"}' http://127.0.0.1:8080/create-student
func createStudent(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("createStudent() - error")
        return
    }

    var newStudent Student
    json.Unmarshal(body, &newStudent)
    Students = append(Students, newStudent)
    json.NewEncoder(w).Encode(newStudent)
    fmt.Println("createStudent() - created new Student")
}

func main() {
    Students = []Student {
        { Id: "0", Name: "Student A", Age: 11, FavouriteSubject: "Maths" },
        { Id: "1", Name: "Student B", Age: 14, FavouriteSubject: "Geography" },
        { Id: "2", Name: "Student C", Age: 8, FavouriteSubject: "English" },
    }
	handleRequests()
}

