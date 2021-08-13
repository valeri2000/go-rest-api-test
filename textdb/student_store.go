package textdb

import (
	"encoding/json"
	"errors"

	db "github.com/valeri2000/go-text-db"
	"main.go/student"
)

type Repo struct {
	Database *db.Database
}

func NewRepo(fileName string) (*Repo, error) {
	tempDatabase, err := db.NewDatabase(fileName)

	if err != nil {
		return nil, err
	}
	return &Repo{Database: tempDatabase}, nil
}

func (repo *Repo) GetStudents() ([]student.Student, error) {
	tempStudents, err := repo.Database.GetAll()
	if err != nil {
		return nil, err
	}
	students := make([]student.Student, len(tempStudents))
	for i := range tempStudents {
		tempStudent := tempStudents[i]

		jsonString, err := json.Marshal(tempStudent)
		if err != nil {
			return nil, errors.New("failed getting all students")
		}

		err = json.Unmarshal(jsonString, &students[i])
		if err != nil {
			return nil, errors.New("failed getting all students")
		}
	}

	return students, err
}

func (repo *Repo) GetStudentByID(id string) (student.Student, error) {
	temp, ok := repo.Database.Get(id)
	if !ok {
		return student.Student{}, errors.New("couldn't get student id")
	}

	stud, ok := temp.(student.Student)
	if !ok {
		return student.Student{}, errors.New("couldn't get student id")
	}
	return stud, nil
}

func (repo *Repo) CountStudents() int {
	return repo.Database.GetCount()
}

func (repo *Repo) HasStudentWithID(id string) bool {
	_, ok := repo.Database.Get(id)
	return ok
}

func (repo *Repo) AddStudent(student student.Student) error {
	return repo.Database.Put(student.Id, student)
}

func (repo *Repo) DeleteStudentByID(id string) error {
	return repo.Database.Put(id, nil)
}

func (repo *Repo) UpdateStudentByID(id string, updatedStudent student.Student) error {
	return repo.Database.Put(id, updatedStudent)
}
