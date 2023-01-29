//http.ServeMux type has ServeHTTP method

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}
type Student struct {
	Name   string
	Age    int
	Rollno int
}

type personHandler struct {
	db *sql.DB
}

type studentHandler struct {
	db *sql.DB
}

type pingHandler string

func openConnectionToDb(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	return db, err
}
func (p personHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rows, err1 := p.db.Query("SELECT * FROM person")
	if err1 != nil {
		log.Fatal(err1)
	}
	var persons []Person
	for rows.Next() {
		var per Person
		err2 := rows.Scan(&per.Name, &per.Age, &per.Phone)
		if err2 != nil {
			log.Fatal(err2)
		}
		persons = append(persons, per)
	}
	a, _ := json.Marshal(persons)
	w.Write(a)
}

func (s studentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rows, err1 := s.db.Query("SELECT * FROM student")
	if err1 != nil {
		log.Fatal(err1)
	}
	var students []Student
	for rows.Next() {
		var stu Student
		err2 := rows.Scan(&stu.Name, &stu.Age, &stu.Rollno)
		if err2 != nil {
			log.Fatal(err2)
		}
		students = append(students, stu)
	}
	a, _ := json.Marshal(students)
	w.Write(a)

}

func (p pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(p))
}

func main() {

	db, err := openConnectionToDb("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")
	if err != nil {
		log.Fatal(err)
	}
	per := personHandler{db}
	stu := studentHandler{db}
	p := pingHandler("pong")

	mux := http.NewServeMux()
	mux.Handle("/person", per)
	mux.Handle("/student", stu)
	mux.Handle("/ping", p)

	http.ListenAndServe(":8005", mux)

}
