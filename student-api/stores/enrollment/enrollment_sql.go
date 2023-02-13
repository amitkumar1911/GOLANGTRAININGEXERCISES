package enrollment

import (
	"database/sql"
	"errors"
)

type EnrollDb struct {
	D *sql.DB
}

func (e EnrollDb) Insert(rollno int, id int) error {
	_, err := e.D.Exec(`INSERT INTO studentSubject VALUES(?,?)`, rollno, id)

	if err != nil {
		return err
	}
	return nil
}

func (e EnrollDb) FindIdByRoll(rollno int) ([]int, error) {
	rows, err := e.D.Query(`SELECT id FROM studentSubject WHERE rollno=?`, rollno)

	if err != nil {
		return []int{}, errors.New("some error occured")
	}

	var r1 []int
	for rows.Next() {
		var r2 int
		err := rows.Scan(&r2)

		if err != nil {
			return []int{}, err
		}

		r1 = append(r1, r2)
	}

	return r1, nil

}
