package main

import (
	"fmt"

	"main.go/rest"
	"main.go/student"
	"main.go/textdb"
)

func main() {
	myRepo, err := textdb.NewRepo("data.json")
	if err != nil {
		fmt.Println("Error while creating repo!")
		return
	}

	myStudentService := &(student.Service{Repo: myRepo})
	server := rest.CreateServer(myStudentService)
	server.Serve()
}
