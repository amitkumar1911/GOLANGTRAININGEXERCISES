package subject

import (
	"errors"
	"fmt"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
)

type subjectstore interface {
	CreateSubject(models.Subject) error
	GetSubject(id int) ([]byte, error)
	CheckSubjectExist(int) bool
	FindNamesById([]int) ([]byte, error)
}

type subjectService struct {
	subStr subjectstore
}

func NewSubjectService(subStr subjectstore) subjectService {
	return subjectService{subStr}
}

func (s subjectService) CreateSubject(sub models.Subject) error {
	if sub.Id == 0 || sub.Name == "" {

		return errors.New("invalid request")

	} else {
		fmt.Println(2)
		err := s.subStr.CreateSubject(sub)
		return err
	}
}

func (s subjectService) GetSubject(id int) ([]byte, error) {
	if id == 0 {

		return []byte{}, errors.New("invalid request")
	}
	return s.subStr.GetSubject(id)
}

func (s subjectService) SubjectExist(id int) bool {

	return s.subStr.CheckSubjectExist(id)
}

func (s subjectService) FindNamesById(id []int) ([]byte, error) {

	return s.subStr.FindNamesById(id)
}
