package memdb

import "main.go/student"

func InitializeService() (service *student.Service) {
	myRepo := &(student.Repo{})
	service = &(student.Service{Repo: myRepo})

	service.AddStudent(student.Student{Id: "0", Name: "Student A", Age: 11, FavouriteSubject: "Maths"})
	service.AddStudent(student.Student{Id: "1", Name: "Student B", Age: 14, FavouriteSubject: "Geography"})
	service.AddStudent(student.Student{Id: "2", Name: "Student C", Age: 8, FavouriteSubject: "English"})
	service.AddStudent(student.Student{Id: "3", Name: "Student D", Age: 9, FavouriteSubject: "History"})

	return service
}
