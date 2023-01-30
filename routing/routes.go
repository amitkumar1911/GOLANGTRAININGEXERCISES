package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Name   string
	Rollno int
	Age    int
}
type mydb struct {
	db *sql.DB
}

func (m mydb) getStudents(w http.ResponseWriter, r *http.Request) {

	rows, err := m.db.Query("SELECT * FROM student")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var stu []student
	for rows.Next() {

		var s student
		err1 := rows.Scan(&s.Name, &s.Rollno, &s.Age)
		if err1 != nil {
			log.Fatal(err1)
		}
		stu = append(stu, s)
	}
	a, err2 := json.Marshal(stu)
	if err2 != nil {
		log.Fatal(err2)
	}
	w.Write(a)
}

func (m mydb) getStudentsByRoll(w http.ResponseWriter, r *http.Request) {

	rno := r.URL.Query().Get("rollno")
	value, _ := strconv.Atoi(rno)
	query := fmt.Sprintf("SELECT * FROM student WHERE rollno=%d", value)

	rows, err := m.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var stu []student
	for rows.Next() {

		var s student
		err1 := rows.Scan(&s.Name, &s.Rollno, &s.Age)
		if err1 != nil {
			log.Fatal(err1)
		}
		stu = append(stu, s)
	}
	a, err2 := json.Marshal(stu)
	if err2 != nil {
		log.Fatal(err2)
	}
	w.Write(a)
}

func (m mydb) insertToDb(s []student) {

	for i := 0; i < len(s); i++ {
		query := fmt.Sprintf(`INSERT INTO student VALUES("%s","%d","%d")`, s[i].Name, s[i].Rollno, s[i].Age)
		m.db.Exec(query)
	}
}

func (m mydb) postStudents(w http.ResponseWriter, r *http.Request) {

	var s []student
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(bytes, &s)
	m.insertToDb(s)
	w.Write([]byte("successfull"))

}

func (m mydb) updateStudents(w http.ResponseWriter, r *http.Request) {

}

func main() {

	conn, err := sql.Open("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {
		log.Fatal(err)
	}

	db := mydb{conn}
	mux := http.NewServeMux()

	h1 := http.HandlerFunc(db.getStudents)
	h2 := http.HandlerFunc(db.getStudentsByRoll)
	h3 := http.HandlerFunc(db.postStudents)

	mux.Handle("/students", h1)
	mux.Handle("/rollno", h2)
	mux.Handle("/post/students", h3)

	http.ListenAndServe(":8007", mux)

}
