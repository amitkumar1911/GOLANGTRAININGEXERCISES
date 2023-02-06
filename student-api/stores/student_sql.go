package stores

import (
	"database/sql"
	"encoding/json"

	"github.com/student-api/models"
)

type SqlDb struct {
	D *sql.DB
}

func (m SqlDb) GetAll() ([]byte, error) {
	rows, err := m.D.Query("SELECT * FROM student")

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
func (m SqlDb) Insert(s models.Student) error {
	_, err := m.D.Exec(`INSERT INTO student VALUES(?,?,?)`, s.Name, s.Rollno, s.Age)

	if err != nil {
		return err
	}
	return nil
}

func (m SqlDb) Get(rollno int) ([]byte, error) {

	rows, err := m.D.Query(`SELECT * FROM student WHERE rollno=?`, rollno)
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
	a, _:= json.Marshal(stu)
	return a, nil

}

func (m SqlDb) Update(rno int, s models.Student) error {

	_, err := m.D.Exec(`UPDATE student SET name=?,age=? WHERE  rollno=?`, s.Name, s.Age, rno)
	if err != nil {
		return err
	}

	return nil
}
