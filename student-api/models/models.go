package models

type Student struct {
	Name   string
	Rollno int
	Age    int
}

type Subject struct {
	Name string
	Id   int
}

type StudentSubject struct {
	Rollno int
	Id     int
}
