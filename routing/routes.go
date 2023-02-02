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
	"github.com/gorilla/mux"
)

type student struct {
	Name   string
	Rollno int
	Age    int
}
type mydb struct {
	db *sql.DB
}

func (m mydb) insertToDb(s []student) {

	for i := 0; i < len(s); i++ {
		query := fmt.Sprintf(`INSERT INTO student VALUES("%s","%d","%d")`, s[i].Name, s[i].Rollno, s[i].Age)
		m.db.Exec(query)
	}
}

func (m mydb) processStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var s []student
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("test error"))
			return
		}
		err1 := json.Unmarshal(bytes, &s)
		if err1 != nil {
			w.Write([]byte("some error occured"))
			return
		}
		m.insertToDb(s)
		w.Write([]byte("data is entered successfully"))

	} else {
		rows, err := m.db.Query("SELECT * FROM student")

		if err != nil {
			w.Write([]byte("something went wrong"))
			return
		}

		defer rows.Close()
		var stu []student
		for rows.Next() {

			var s student
			err1 := rows.Scan(&s.Name, &s.Rollno, &s.Age)
			if err1 != nil {
				w.Write([]byte("cannot process rows"))
				return
			}
			stu = append(stu, s)
		}
		a, _ := json.Marshal(stu)
		w.Write(a)

	}

}

func (m mydb) filterByRoll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		var s []student
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("test error"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err1 := json.Unmarshal(bytes, &s)
		if err1 != nil {
			w.Write([]byte("some error occured"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		params := mux.Vars(r)
		value, _ := strconv.Atoi(params["id"])
		query := fmt.Sprintf(`UPDATE student SET name="%s",rollno="%d",age="%d" WHERE rollno="%d"`, s[0].Name, s[0].Rollno, s[0].Age, value)
		_, err2 := m.db.Exec(query)
		if err2 != nil {
			w.Write([]byte("cannot update data"))
			return
		}
		w.Write([]byte("data is updated successfully"))
	} else {
		params := mux.Vars(r)
		value, _ := strconv.Atoi(params["id"])
		query := fmt.Sprintf("SELECT * FROM student WHERE rollno=%d", value)

		rows, err := m.db.Query(query)

		if err != nil {
			w.Write([]byte("something went wrong"))
			return
		}

		defer rows.Close()
		var stu []student
		for rows.Next() {

			var s student
			err1 := rows.Scan(&s.Name, &s.Rollno, &s.Age)
			if err1 != nil {
				w.Write([]byte("cannot process rows"))
				return
			}
			stu = append(stu, s)
		}
		a, _ := json.Marshal(stu)
		w.Write(a)
	}
}

func main() {

	conn, err := sql.Open("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {
		log.Fatal(err)
	}

	db := mydb{conn}
	router := mux.NewRouter()

	router.HandleFunc("/students", db.processStudent)
	router.HandleFunc("/students/roll/{id}", db.filterByRoll)

	http.ListenAndServe(":8007", router)

}
