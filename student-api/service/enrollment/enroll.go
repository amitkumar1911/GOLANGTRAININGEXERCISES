package enrollment

type enrollmentStore interface {
	Insert(int, int) error
	FindRollById(int) ([]int, error)
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

func (enr enrollmentService) FindRollById(id int) ([]int, error) {

	return enr.enStr.FindRollById(id)
}
