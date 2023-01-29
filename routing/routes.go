package routing

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

type mydb struct {
	db *sql.DB
}

func openConnectionToDb(driverName string, dataSourceName string) *sql.DB {
	db, _ := sql.Open(driverName, dataSourceName)
	return db
}

func getPersonDetails(db *sql.DB) ([]Person, error) {
	rows, err1 := db.Query("SELECT * FROM person")
	if err1 != nil {
		return nil, err1
	}
	var persons []Person
	for rows.Next() {
		var per Person
		err2 := rows.Scan(&per.Name, &per.Age, &per.Phone)
		if err2 != nil {
			return nil, err2
		}
		persons = append(persons, per)
	}
	return persons, nil
}

func (conn mydb) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == ("/ping") {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")
	} else if r.URL.Path == "/person" {
		w.WriteHeader(http.StatusOK)
		value, _ := getPersonDetails(conn.db)
		a, _ := json.Marshal(value)
		w.Write(a)
	} else {
		w.Write([]byte("invalid"))
	}
}

func main() {
	conn := openConnectionToDb("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	db := mydb{conn}
	// insertToPerson(conn.DB)
	http.ListenAndServe(":8000", http.HandlerFunc(db.rootHandler))
}
