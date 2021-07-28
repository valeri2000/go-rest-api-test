package main

import (
	"main.go/memdb"
	"main.go/rest"
)

func main() {
	myStudentService := memdb.InitializeService()
	server := rest.CreateServer(myStudentService)
	server.Serve()
}
