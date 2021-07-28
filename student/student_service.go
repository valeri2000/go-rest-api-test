package student

import "errors"

type Service struct {
	Repo *Repo
}

func (service *Service) GetStudents() []Student {
	return service.Repo.GetStudents()
}

func (service *Service) GetStudentByID(id string) (Student, error) {
	if service.Repo.HasStudentWithID(id) == false {
		return Student{}, errors.New("Invalid ID for deletion!")
	}

	return service.Repo.GetStudentByID(id), nil
}

func (service *Service) CountStudents() int {
	return service.Repo.CountStudents()
}

func (service *Service) AddStudent(student Student) error {
	service.Repo.AddStudent(student)
	return nil
}

func (service *Service) DeleteStudentByID(id string) error {
	if service.Repo.HasStudentWithID(id) == false {
		return errors.New("Invalid ID for deletion!")
	}

	service.Repo.DeleteStudentByID(id)
	return nil
}

func (service *Service) UpdateStudentByID(id string, updatedStudent Student) error {
	if service.Repo.HasStudentWithID(id) == false {
		return errors.New("Invalid ID for deletion!")
	}

	service.Repo.UpdateStudentByID(id, updatedStudent)
	return nil
}
