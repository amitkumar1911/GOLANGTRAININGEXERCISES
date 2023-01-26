package routing

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	Name  string
	Age   int
	Phone string
}

type apiHandler string

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

func (a apiHandler) getPerson(w http.ResponseWriter, r *http.Request) error {

	db, err := openConnectionToDb("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {
		return err
	}

	rows, err1 := db.Query("SELECT * FROM person")

	if err1 != nil {
		return err1
	}

	var persons []person

	for rows.Next() {
		var per person

		err = rows.Scan(&per.Name, &per.Age, &per.Phone)

		persons = append(persons, per)
	}

	fmt.Println(persons)
	return nil
}

func (a apiHandler) insertToPerson(w http.ResponseWriter, r *http.Request) error {

	db, err := openConnectionToDb("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	if err != nil {
		return err
	}

	_, err1 := db.Exec("INSERT INTO person VALUES(`amit`,`23`,`123`),(`aman`,`24`,`345`),(`akash`,25,`678`)")

	if err1 != nil {
		return err1
	}
	return nil
}
