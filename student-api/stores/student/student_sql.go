package student

import (
	"database/sql"
	"encoding/json"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
)

type StudentDb struct {
	D *sql.DB
}

func (m StudentDb) Get(rollno int) ([]byte, error) {

	var rows *sql.Rows
	var err error

	rows, err = m.D.Query(`SELECT * FROM student where rollno=?`, rollno)

	if err != nil {

		return nil, err
	}

	defer rows.Close()
	var stu []models.Student
	for rows.Next() {

		var s models.Student
		err1 := rows.Scan(&s.Name, &s.Rollno, &s.Age)
		if err1 != nil {
			return nil, err1
		}
		stu = append(stu, s)
	}
	a, _ := json.Marshal(stu)
	return a, nil
}
func (m StudentDb) Create(s models.Student) error {

	_, err := m.D.Exec(`INSERT INTO student VALUES(?,?,?)`, s.Name, s.Rollno, s.Age)

	if err != nil {
		return err
	}
	return nil

}

func (m StudentDb) StudentExist(rollno int) bool {

	count := 0
	rows := m.D.QueryRow(`SELECT COUNT(*) FROM student WHERE rollno=?`, rollno)

	err := rows.Scan(&count)

	if err != nil {
		return false
	}
	return count != 0
}
