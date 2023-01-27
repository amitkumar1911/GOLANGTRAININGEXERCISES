package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type testdb struct {
	Name   string
	Age    int
	Rollno int
}

func openConnection(mydriver string, connstring string) (*sql.DB, error) {

	db, err := sql.Open(mydriver, connstring)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func insert(db *sql.DB) error {

	_, err := db.Exec(`INSERT INTO testdb VALUES('a',21,1),('b',22,2)`)

	if err != nil {

		return err

	}

	fmt.Println("inserted successfully")
	return nil
}

func retrive(db *sql.DB) ([]testdb, error) {

	rows, err := db.Query("SELECT * FROM testdb")

	if err != nil {
		return nil, err
	}

	var test []testdb
	for rows.Next() {

		var t testdb

		rows.Scan(&t.Name, &t.Age, &t.Rollno)

		test = append(test, t)

	}
	return test, nil
}

func WriteFile(file *os.File, t []testdb) error {

	for i := 0; i < len(t); i++ {

		a := strconv.Itoa(t[i].Age)
		b := strconv.Itoa(t[i].Rollno)

		file.Write([]byte(t[i].Name))
		file.Write([]byte(","))
		file.Write([]byte(a))
		file.Write([]byte(","))
		file.Write([]byte(b))

	}
	return nil
}

func main() {

	db, _ := openConnection("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	// insert(db)

	value, _ := retrive(db)

	file, _ := os.Create("temp.txt")

	WriteFile(file, value)

}
