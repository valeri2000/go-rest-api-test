package memdb

import "main.go/student"

type Repo struct {
	Students []student.Student `json:"students"`
}

func (repo *Repo) GetStudents() []student.Student {
	data := make([]student.Student, repo.CountStudents())
	copy(data, repo.Students)
	return data
}

func (repo *Repo) GetStudentByID(id string) student.Student {
	for _, student := range repo.Students {
		if student.Id == id {
			return student
		}
	}
	return student.Student{}
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

func (repo *Repo) AddStudent(student student.Student) {
	repo.Students = append(repo.Students, student)
}

func (repo *Repo) DeleteStudentByID(id string) {
	for i, student := range repo.Students {
		if student.Id == id {
			repo.Students = append(repo.Students[:i], repo.Students[i+1:]...)
			break
		}
	}
}

func (repo *Repo) UpdateStudentByID(id string, updatedStudent student.Student) {
	for i, student := range repo.Students {
		if student.Id == id {
			repo.Students[i] = updatedStudent
			break
		}
	}
}
