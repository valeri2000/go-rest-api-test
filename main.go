package main

import (
	"fmt"
    "strconv"
	"log"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
)

type Student struct {
    Id string `json:"Id"`
    Name string `json:"Name"`
    Age int `json:"Age"`
    FavouriteSubject string `json:"FavouriteSubject"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Number of students -", strconv.Itoa(len(Students)))
    fmt.Println("homePage()")
}

func handleRequests() {
    Router := mux.NewRouter().StrictSlash(true)
    Router.HandleFunc("/", homePage)
    Router.HandleFunc("/list-students", getStudents)
    Router.HandleFunc("/create-student", createStudent).Methods("POST")
    Router.HandleFunc("/view-student/{id}", getOneStudent)
    Router.HandleFunc("/delete-student/{id}", deleteStudent).Methods("DELETE")
    Router.HandleFunc("/update-student/{id}", updateStudent).Methods("PUT")

    log.Fatal(http.ListenAndServe(":8080", Router))
}

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
    fmt.Println("createStudent()")

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

// via DELETE
// curl -X "DELETE" http://127.0.0.1:8080/delete-student/0
func deleteStudent(w http.ResponseWriter, r *http.Request) {
    fmt.Println("deleteStudent()")

    vars := mux.Vars(r)
    fmt.Println("mux.Vars():", vars)

    requestedId := vars["id"]
    for i := range Students {
        if Students[i].Id == requestedId {
            json.NewEncoder(w).Encode(Students[i])
            Students = append(Students[:i], Students[i+1:]...)
            break
        }
    }
}

// via PUT
// curl -X "PUT" -H 'Content-Type: application/json' -d '{"Id": "0", "Name": "Student AA", "Age": 11, "FavouriteSubject": "Maths"}' http://127.0.0.1:8080/update-student/0
func updateStudent(w http.ResponseWriter, r *http.Request) {
    fmt.Println("updateStudent()")

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("updateStudent() - error")
        return
    }

    var updatedStudent Student
    json.Unmarshal(body, &updatedStudent)

    vars := mux.Vars(r)
    fmt.Println("mux.Vars():", vars)

    requestedId := vars["id"]
    for i := range Students {
        if Students[i].Id == requestedId {
            json.NewEncoder(w).Encode(updatedStudent)
            Students[i] = updatedStudent
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

