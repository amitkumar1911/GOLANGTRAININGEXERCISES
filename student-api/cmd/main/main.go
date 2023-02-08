package main

import (
	"database/sql"
	"net/http"

	"github.com/GOLANGTRAININGEXERCISES/student-api/handler"
	enrollService "github.com/GOLANGTRAININGEXERCISES/student-api/service/enrollment"
	stuService "github.com/GOLANGTRAININGEXERCISES/student-api/service/student"
	subService "github.com/GOLANGTRAININGEXERCISES/student-api/service/subject"

	enrollDb "github.com/GOLANGTRAININGEXERCISES/student-api/stores/enrollment"
	stuDb "github.com/GOLANGTRAININGEXERCISES/student-api/stores/student"
	subDb "github.com/GOLANGTRAININGEXERCISES/student-api/stores/subject"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	conn, _ := sql.Open("mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings")

	studentStore := stuDb.StudentDb{conn}
	subjectStore := subDb.SubjectDb{conn}
	enrollmentStore := enrollDb.EnrollDb{conn}

	subjectService := subService.NewSubjectService(subjectStore)
	enrollmentService := enrollService.NewEnrollmentService(enrollmentStore)
	studentService := stuService.NewStudentService(studentStore, subjectService, enrollmentService)

	studentHandler := handler.NewStudentHandler(studentService)
	subjectHandler := handler.NewSubjectHandler(subjectService)

	router := mux.NewRouter()
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")
	router.HandleFunc("/students/{rollno}", studentHandler.Get).Methods("GET")

	router.HandleFunc("/subjects", subjectHandler.Create).Methods("POST")
	router.HandleFunc("/subjects/{id}", subjectHandler.Get).Methods("GET")

	router.HandleFunc("/students/{rollno}/subjects/{id}", studentHandler.EnrollSubject).Methods("POST")

	http.ListenAndServe(":8080", router)
}
