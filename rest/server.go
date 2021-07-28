package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/student"
)

type Server struct {
	StudentService *student.Service
	Router         *mux.Router
}

func CreateServer(newStudentService *student.Service) *Server {
	return &Server{
		StudentService: newStudentService,
		Router:         mux.NewRouter().StrictSlash(true),
	}
}

func (server *Server) Serve() {
	server.Router.HandleFunc("/", server.homePage)
	server.Router.HandleFunc("/students", server.getStudents).Methods("GET")
	server.Router.HandleFunc("/students", server.createStudent).Methods("POST")
	server.Router.HandleFunc("/student/{id}", server.getOneStudent).Methods("GET")
	server.Router.HandleFunc("/student/{id}", server.deleteStudent).Methods("DELETE")
	server.Router.HandleFunc("/student/{id}", server.updateStudent).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
