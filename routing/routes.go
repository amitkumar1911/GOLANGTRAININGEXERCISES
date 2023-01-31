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

	if r.Method != "GET" {
		w.Write([]byte("expected get found some other method"))
	}

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

	if r.Method != "POST" {
		w.Write([]byte("expected post found some other method"))
	} else {
		var s []student
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("cannot unmarshal the data"))
			return
		}
		err1 := json.Unmarshal(bytes, &s)
		if err1 != nil {
			w.Write([]byte("some error occured"))
			return
		}
		m.insertToDb(s)
		w.Write([]byte("data is entered successfully"))

	}

}

func (m mydb) updateStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Write([]byte("expected Put found some other method"))
	} else {

		var s []student
		fmt.Println(r.Body)
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("cannot convert to slice of bytes"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err1 := json.Unmarshal(bytes, &s)
		if err1 != nil {
			w.Write([]byte("some error occured"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		rno := r.URL.Query().Get("rollno")
		value, _ := strconv.Atoi(rno)
		query := fmt.Sprintf("UPDATE student SET name=%s,rollno=%d,age=%d WHERE rollno=%d", s[0].Name, s[0].Rollno, s[0].Age, value)
		_, err2 := m.db.Exec(query)
		if err2 != nil {
			w.Write([]byte("cannot update data"))
			return
		}
		w.Write([]byte("data is updated successfully"))
	}
}

func main() {

	conn, err := sql.Open("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {
		log.Fatal(err)
	}

	db := mydb{conn}
	mux := http.NewServeMux()

	// h1 := http.HandlerFunc(db.getStudents)
	// h2 := http.HandlerFunc(db.getStudentsByRoll)
	// h3 := http.HandlerFunc(db.postStudents)
	// h4 := http.HandlerFunc(db.updateStudents)

	mux.HandleFunc("/students", db.getStudents)
	mux.HandleFunc("/rollno", db.getStudentsByRoll)
	mux.HandleFunc("/post/student", db.postStudents)
	mux.HandleFunc("/update/student", db.updateStudents)

	http.ListenAndServe(":8007", mux)

}
