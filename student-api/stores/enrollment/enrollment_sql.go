package enrollment

import (
	"database/sql"
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
