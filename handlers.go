package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (restServer *RestServer) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Number of students -", strconv.Itoa(len(restServer.Students)))
	fmt.Println("homePage()")
}

func (restServer *RestServer) getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getStudents()")
	json.NewEncoder(w).Encode(restServer.Students)
}

func (restServer *RestServer) getOneStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOneStudent()")

	vars := mux.Vars(r)
	fmt.Println("mux.Vars():", vars)

	requestedId := vars["id"]
	for i := range restServer.Students {
		if restServer.Students[i].Id == requestedId {
			json.NewEncoder(w).Encode(restServer.Students[i])
		}
	}
}

// via POST
// curl -X POST -H 'Content-Type: application/json' -d '{"id": "3", "name": "Student D", "age": 12, "favouriteSubject": "History"}' http://127.0.0.1:8080/students
func (restServer *RestServer) createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createStudent()")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("createStudent() - error")
		return
	}

	var newStudent Student
	json.Unmarshal(body, &newStudent)
	restServer.Students = append(restServer.Students, newStudent)
	json.NewEncoder(w).Encode(newStudent)
	fmt.Println("createStudent() - created new Student")
}

// via DELETE
// curl -X "DELETE" http://127.0.0.1:8080/student/0
func (restServer *RestServer) deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteStudent()")

	vars := mux.Vars(r)
	fmt.Println("mux.Vars():", vars)

	requestedId := vars["id"]
	for i := range restServer.Students {
		if restServer.Students[i].Id == requestedId {
			json.NewEncoder(w).Encode(restServer.Students[i])
			restServer.Students = append(restServer.Students[:i], restServer.Students[i+1:]...)
			break
		}
	}
}

// via PUT
// curl -X "PUT" -H 'Content-Type: application/json' -d '{"id": "0", "name": "Student AA", "age": 11, "favouriteSubject": "Maths"}' http://127.0.0.1:8080/student/0
func (restServer *RestServer) updateStudent(w http.ResponseWriter, r *http.Request) {
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
	for i := range restServer.Students {
		if restServer.Students[i].Id == requestedId {
			json.NewEncoder(w).Encode(updatedStudent)
			restServer.Students[i] = updatedStudent
		}
	}
}
