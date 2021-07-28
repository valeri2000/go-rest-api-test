package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Students []Student
	Router   *mux.Router
}

func initApp(students []Student) *App {
	return &App{
		Students: students,
		Router:   mux.NewRouter().StrictSlash(true),
	}
}

func (app *App) serve() {
	app.Router.HandleFunc("/", app.homePage)
	app.Router.HandleFunc("/students", app.getStudents).Methods("GET")
	app.Router.HandleFunc("/students", app.createStudent).Methods("POST")
	app.Router.HandleFunc("/student/{id}", app.getOneStudent).Methods("GET")
	app.Router.HandleFunc("/student/{id}", app.deleteStudent).Methods("DELETE")
	app.Router.HandleFunc("/student/{id}", app.updateStudent).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
