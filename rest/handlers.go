package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"main.go/student"
)

func (server *Server) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("[15:04:05]"), "homePage()")
	numberOfStudents := server.StudentService.CountStudents()
	fmt.Fprintln(w, "Number of students: ", strconv.Itoa(numberOfStudents))
}

func (server *Server) getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("[15:04:05]"), "getStudents()")
	data := server.StudentService.GetStudents()
	json.NewEncoder(w).Encode(data)
}

func (server *Server) getOneStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("[15:04:05]"), "getOneStudent()")

	vars := mux.Vars(r)

	requestedId := vars["id"]
	requestedStudent, err := server.StudentService.GetStudentByID(requestedId)

	if err != nil {
		fmt.Println(time.Now().Format("[15:04:05]"), err)
		return
	}

	json.NewEncoder(w).Encode(requestedStudent)
}

// via POST
// curl -X POST -H 'Content-Type: application/json' -d '{"id": "3", "name": "Student D", "age": 12, "favouriteSubject": "History"}' http://127.0.0.1:8080/students
func (server *Server) createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("[15:04:05]"), "createStudent()")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(time.Now().Format("[15:04:05]"), err)
		return
	}

	var newStudent student.Student
	json.Unmarshal(body, &newStudent)
	server.StudentService.AddStudent(newStudent)
	json.NewEncoder(w).Encode(newStudent)
}

// via DELETE
// curl -X "DELETE" http://127.0.0.1:8080/student/0
func (server *Server) deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("[15:04:05]"), "deleteStudent()")

	vars := mux.Vars(r)

	requestedId := vars["id"]
	err := server.StudentService.DeleteStudentByID(requestedId)

	if err != nil {
		fmt.Println(time.Now().Format("[15:04:05]"), err)
		return
	}
}

// via PUT
// curl -X "PUT" -H 'Content-Type: application/json' -d '{"id": "0", "name": "Student AA", "age": 11, "favouriteSubject": "Maths"}' http://127.0.0.1:8080/student/0
func (server *Server) updateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("[15:04:05]"), "updateStudent()")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(time.Now().Format("[15:04:05]"), err)
		return
	}

	var updatedStudent student.Student
	json.Unmarshal(body, &updatedStudent)

	vars := mux.Vars(r)

	requestedId := vars["id"]
	err = server.StudentService.UpdateStudentByID(requestedId, updatedStudent)

	if err != nil {
		fmt.Println(time.Now().Format("[15:04:05]"), err)
		return
	}
}
