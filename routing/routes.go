package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

func openConnectionToDb(driverName string, dataSourceName string) (*sql.DB, error) {

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {

		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(err.Error())
	}

	return db, nil

}

func getPersonDetails(db *sql.DB) ([]Person, error) {

	count := 0

	rows, err1 := db.Query("SELECT * FROM person")

	if err1 != nil {

		return nil, err1
	}

	var persons []Person

	for rows.Next() {
		count++
		var per Person

		err2 := rows.Scan(&per.Name, &per.Age, &per.Phone)

		if err2 != nil {
			return nil, err2
		}

		persons = append(persons, per)
		if count == 1 {
			break
		}
	}
	return persons, nil
}

func insertToPerson() ([]Person, error) {

	db, err := openConnectionToDb("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {

		return nil, err
	}

	p := Person{"amit", 21, "123"}

	query := fmt.Sprintf(`INSERT INTO person VALUES("%s","%d","%s")`, p.Name, p.Age, p.Phone)

	_, err1 := db.Exec(query)

	if err1 != nil {

		return nil, err1
	}
	value, err2 := getPersonDetails(db)
	return value, err2
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == ("/ping") {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "pong")

	} else if r.URL.Path == "/person" {
		value, _ := insertToPerson()
		w.WriteHeader(http.StatusOK)
		a, _ := json.Marshal(value)
		w.Write(a)

	} else {
		fmt.Fprintln(w, "invalid")
	}
}

func main() {

	// var w http.ResponseWriter
	// var r *http.Request

	err := http.ListenAndServe(":8000", http.HandlerFunc(rootHandler))

	if err != nil {
		fmt.Println("some error occured while starting server")
	}

}
