package main

import (
	"database/sql"
	"net/http"

	stuHandler "github.com/GOLANGTRAININGEXERCISES/student-api/handler/student"
	subHandler "github.com/GOLANGTRAININGEXERCISES/student-api/handler/subject"

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

	studentHandler := stuHandler.NewStudentHandler(studentService)
	subjectHandler := subHandler.NewSubjectHandler(subjectService)

	router := mux.NewRouter()
	router.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{rollno}", studentHandler.GetStudent).Methods("GET")

	router.HandleFunc("/subjects", subjectHandler.CreateSubject).Methods("POST")
	router.HandleFunc("/subjects/{id}", subjectHandler.GetSubject).Methods("GET")

	router.HandleFunc("/students/{rollno}/subjects/{id}", studentHandler.EnrollStudent).Methods("POST")
	router.HandleFunc("/students/{rollno}/subjects", studentHandler.GetNames).Methods("GET")

	http.ListenAndServe(":8080", router)
}
