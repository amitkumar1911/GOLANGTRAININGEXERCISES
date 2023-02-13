package enrollment

type enrollmentStore interface {
	Insert(int, int) error
	FindIdByRoll(int) ([]int, error)
}

type enrollmentService struct {
	enStr enrollmentStore
}

func NewEnrollmentService(enr enrollmentStore) enrollmentService {
	return enrollmentService{enr}
}

func (enr enrollmentService) Insert(rollno int, id int) error {

	return enr.enStr.Insert(rollno, id)

}

func (enr enrollmentService) FindIdByRoll(id int) ([]int, error) {

	return enr.enStr.FindIdByRoll(id)
}
