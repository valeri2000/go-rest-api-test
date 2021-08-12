package student

type StudentRepo interface {
	GetStudents() ([]Student, error)
	GetStudentByID(id string) (Student, error)
	CountStudents() int
	HasStudentWithID(id string) bool
	AddStudent(student Student) error
	DeleteStudentByID(id string) error
	UpdateStudentByID(id string, updatedStudent Student) error
}
