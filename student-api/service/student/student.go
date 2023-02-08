package student

import (
	"errors"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
)

type studentstore interface {
	Create(models.Student) error
	Get(int) ([]byte, error)
	StudentExist(int) bool
}

type subjectService interface {
	SubjectExist(int) bool
}

type enrollmentService interface {
	Insert(int, int) error
}

type studentService struct {
	stuStr studentstore
	subSvc subjectService
	enrSvc enrollmentService
}

func NewStudentService(stuStr studentstore, subSvc subjectService, enrSvc enrollmentService) studentService {

	return studentService{stuStr, subSvc, enrSvc}
}
func (s studentService) CreateStudent(stu models.Student) error {

	if stu.Name == "" || stu.Rollno == 0 {

		return errors.New("invalid request")

	} else {
		err := s.stuStr.Create(stu)
		return err
	}

}

func (s studentService) GetStudent(rollno int) ([]byte, error) {

	if rollno == 0 {

		return []byte{}, errors.New("invalid request")
	}
	return s.stuStr.Get(rollno)
}

func (s studentService) CheckExist(rollno int, id int) error {

	value1 := s.stuStr.StudentExist(rollno)

	if !value1 {
		return errors.New("invalid params")
	}
	value2 := s.subSvc.SubjectExist(id)

	if !value2 {
		return errors.New("invalid params")
	}
	return s.enrSvc.Insert(rollno, id)
}
