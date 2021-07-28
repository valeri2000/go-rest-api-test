package main

func main() {
	restServer := initServer([]Student{
		{Id: "0", Name: "Student A", Age: 11, FavouriteSubject: "Maths"},
		{Id: "1", Name: "Student B", Age: 14, FavouriteSubject: "Geography"},
		{Id: "2", Name: "Student C", Age: 8, FavouriteSubject: "English"},
	})

	restServer.serve()
}
