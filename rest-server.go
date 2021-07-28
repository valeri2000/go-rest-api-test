package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RestServer struct {
	Students []Student
	Router   *mux.Router
}

func initServer(students []Student) *RestServer {
	return &RestServer{
		Students: students,
		Router:   mux.NewRouter().StrictSlash(true),
	}
}

func (restServer *RestServer) serve() {
	restServer.Router.HandleFunc("/", restServer.homePage)
	restServer.Router.HandleFunc("/students", restServer.getStudents).Methods("GET")
	restServer.Router.HandleFunc("/students", restServer.createStudent).Methods("POST")
	restServer.Router.HandleFunc("/student/{id}", restServer.getOneStudent).Methods("GET")
	restServer.Router.HandleFunc("/student/{id}", restServer.deleteStudent).Methods("DELETE")
	restServer.Router.HandleFunc("/student/{id}", restServer.updateStudent).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", restServer.Router))
}
