package memdb

import (
	"errors"

	"main.go/student"
)

type Repo struct {
	Students []student.Student `json:"students"`
}

func (repo *Repo) GetStudents() ([]student.Student, error) {
	data := make([]student.Student, repo.CountStudents())
	copy(data, repo.Students)
	return data, nil
}

func (repo *Repo) GetStudentByID(id string) (student.Student, error) {
	for _, student := range repo.Students {
		if student.Id == id {
			return student, nil
		}
	}
	return student.Student{}, errors.New("couldn't get student id")
}

func (repo *Repo) CountStudents() int {
	return len(repo.Students)
}

func (repo *Repo) HasStudentWithID(id string) bool {
	for _, student := range repo.Students {
		if student.Id == id {
			return true
		}
	}
	return false
}

func (repo *Repo) AddStudent(student student.Student) error {
	repo.Students = append(repo.Students, student)
	return nil
}

func (repo *Repo) DeleteStudentByID(id string) error {
	ok := false
	for i, student := range repo.Students {
		if student.Id == id {
			repo.Students = append(repo.Students[:i], repo.Students[i+1:]...)
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("couldn't delete student")
	}
	return nil
}

func (repo *Repo) UpdateStudentByID(id string, updatedStudent student.Student) error {
	ok := false
	for i, student := range repo.Students {
		if student.Id == id {
			repo.Students[i] = updatedStudent
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("couldn't update student")
	}
	return nil
}
