package student

import "errors"

type Service struct {
	Repo StudentRepo
}

func (service *Service) GetStudents() ([]Student, error) {
	return service.Repo.GetStudents()
}

func (service *Service) GetStudentByID(id string) (Student, error) {
	if service.Repo.HasStudentWithID(id) == false {
		return Student{}, errors.New("Invalid ID for deletion!")
	}

	return service.Repo.GetStudentByID(id)
}

func (service *Service) CountStudents() int {
	return service.Repo.CountStudents()
}

func (service *Service) AddStudent(student Student) error {
	return service.Repo.AddStudent(student)
}

func (service *Service) DeleteStudentByID(id string) error {
	if service.Repo.HasStudentWithID(id) == false {
		return errors.New("Invalid ID for deletion!")
	}

	return service.Repo.DeleteStudentByID(id)
}

func (service *Service) UpdateStudentByID(id string, updatedStudent Student) error {
	if service.Repo.HasStudentWithID(id) == false {
		return errors.New("Invalid ID for deletion!")
	}

	return service.Repo.UpdateStudentByID(id, updatedStudent)
}
