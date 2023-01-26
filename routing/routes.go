package main

import (
	"database/sql"
	"fmt"
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

func insertToPerson() error {

	db, err := openConnectionToDb("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {

		return err
	}

	query := fmt.Sprintf(`INSERT INTO person`)

	_, err1 := db.Exec("INSERT INTO person (name, age, phone) VALUES(`amit`,23,`123`),(`aman`,24,`345`),(`akash`,25,`678`)")

	if err1 != nil {

		return err1
	}
	getPersonDetails(db)
	return nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == ("/ping") {
		fmt.Fprintln(w, "pong")

	} else {

		value := insertToPerson()
		fmt.Fprintln(w, value)

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
