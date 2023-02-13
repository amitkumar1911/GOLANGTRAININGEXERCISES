package subject

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
)

type SubjectDb struct {
	D *sql.DB
}

func (m SubjectDb) CreateSubject(s models.Subject) error {
	_, err := m.D.Exec(`INSERT INTO subject VALUES(?,?)`, s.Name, s.Id)

	if err != nil {
		return err
	}
	return nil
}

func (m SubjectDb) GetSubject(id int) ([]byte, error) {

	rows, err := m.D.Query(`SELECT * FROM subject WHERE id=?`, id)

	if err != nil {

		return nil, err
	}

	defer rows.Close()
	var sub []models.Subject
	for rows.Next() {

		var s models.Subject
		err1 := rows.Scan(&s.Name, &s.Id)
		if err1 != nil {
			return nil, err1
		}
		sub = append(sub, s)
	}
	a, _ := json.Marshal(sub)
	return a, nil
}

func (m SubjectDb) CheckSubjectExist(id int) bool {

	count := 0
	rows := m.D.QueryRow(`SELECT COUNT(*) FROM subject WHERE id=?`, id)

	err := rows.Scan(&count)
	if err != nil {
		return false
	}

	return count != 0
}

func (m SubjectDb) FindNamesById(r []int) ([]byte, error) {

	var names []string
	for i := 0; i < len(r); i++ {

		var n string
		rows := m.D.QueryRow(`SELECT name from subject WHERE id=?`, r[i])

		err := rows.Scan(&n)

		if err != nil {
			return []byte{}, errors.New("some error occured")
		}

		names = append(names, n)
	}
	a, _ := json.Marshal(names)
	return a, nil
}
