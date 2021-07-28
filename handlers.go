package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *App) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Number of students -", strconv.Itoa(len(app.Students)))
	fmt.Println("homePage()")
}

func (app *App) getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getStudents()")
	json.NewEncoder(w).Encode(app.Students)
}

func (app *App) getOneStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getOneStudent()")

	vars := mux.Vars(r)
	fmt.Println("mux.Vars():", vars)

	requestedId := vars["id"]
	for i := range app.Students {
		if app.Students[i].Id == requestedId {
			json.NewEncoder(w).Encode(app.Students[i])
		}
	}
}

// via POST
// curl -X POST -H 'Content-Type: application/json' -d '{"Id": "3", "Name": "Student D", "Age": 12, "FavouriteSubject": "History"}' http://127.0.0.1:8080/students
func (app *App) createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createStudent()")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("createStudent() - error")
		return
	}

	var newStudent Student
	json.Unmarshal(body, &newStudent)
	app.Students = append(app.Students, newStudent)
	json.NewEncoder(w).Encode(newStudent)
	fmt.Println("createStudent() - created new Student")
}

// via DELETE
// curl -X "DELETE" http://127.0.0.1:8080/student/0
func (app *App) deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteStudent()")

	vars := mux.Vars(r)
	fmt.Println("mux.Vars():", vars)

	requestedId := vars["id"]
	for i := range app.Students {
		if app.Students[i].Id == requestedId {
			json.NewEncoder(w).Encode(app.Students[i])
			app.Students = append(app.Students[:i], app.Students[i+1:]...)
			break
		}
	}
}

// via PUT
// curl -X "PUT" -H 'Content-Type: application/json' -d '{"Id": "0", "Name": "Student AA", "Age": 11, "FavouriteSubject": "Maths"}' http://127.0.0.1:8080/student/0
func (app *App) updateStudent(w http.ResponseWriter, r *http.Request) {
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
	for i := range app.Students {
		if app.Students[i].Id == requestedId {
			json.NewEncoder(w).Encode(updatedStudent)
			app.Students[i] = updatedStudent
		}
	}
}
