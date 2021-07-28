package student

type Repo struct {
	Students []Student `json:"students"`
}

func (repo *Repo) GetStudents() []Student {
	data := make([]Student, repo.CountStudents())
	copy(data, repo.Students)
	return data
}

func (repo *Repo) GetStudentByID(id string) Student {
	for _, student := range repo.Students {
		if student.Id == id {
			return student
		}
	}
	return Student{}
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

func (repo *Repo) AddStudent(student Student) {
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

func (repo *Repo) UpdateStudentByID(id string, updatedStudent Student) {
	for i, student := range repo.Students {
		if student.Id == id {
			repo.Students[i] = updatedStudent
			break
		}
	}
}
