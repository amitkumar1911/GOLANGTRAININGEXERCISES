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
	FindNamesById([]int) ([]byte, error)
}

type enrollmentService interface {
	Insert(int, int) error
	FindIdByRoll(int) ([]int, error)
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

func (s studentService) GetId(rollno int) ([]byte, error) {

	value, err := s.enrSvc.FindIdByRoll(rollno)
	if err != nil {
		return []byte{}, err
	}

	value1, err1 := s.subSvc.FindNamesById(value)

	return value1, err1

}
