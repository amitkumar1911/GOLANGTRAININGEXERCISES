package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/student-api/handler"
	"github.com/student-api/stores"
)

func main() {

	conn, _ := sql.Open("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	studentHandler := handler.NewHandler(stores.SqlDb{conn})

	router := mux.NewRouter()
	router.HandleFunc("/students", studentHandler.GetAll).Methods("GET")
	router.HandleFunc("/students", studentHandler.Insert).Methods("POST")
	router.HandleFunc("/students/roll/{id}", studentHandler.Update).Methods("PUT")
	router.HandleFunc("/students/roll/{id}", studentHandler.Get).Methods("GET")
	http.ListenAndServe(":8080", router)

}
