package student

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
	"github.com/gorilla/mux"
)

type studentService interface {
	CreateStudent(models.Student) error
	GetStudent(int) ([]byte, error)
	CheckExist(int, int) error
	GetId(int) ([]byte, error)
}

type studentHandler struct {
	d studentService
}

func NewStudentHandler(s studentService) studentHandler {
	return studentHandler{s}
}

func (h studentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	value, _ := strconv.Atoi(params["rollno"])

	resp, err := h.d.GetStudent(value)

	if err != nil {
		fmt.Fprint(w, "failed to get all records")
		return
	}
	fmt.Fprint(w, string(resp))
}

func (h studentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {

	var s models.Student
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "cannot convert to slice of bytes")
		return
	}
	err1 := json.Unmarshal(bytes, &s)
	if err1 != nil {
		fmt.Fprint(w, "some error occured while parsing the json data")
		return
	}
	err2 := h.d.CreateStudent(s)
	if err2 != nil {
		fmt.Fprint(w, "cannot insert record")
		return
	}
	fmt.Fprint(w, "inserted data successfully")

}

func (h studentHandler) EnrollStudent(w http.ResponseWriter, r *http.Request) {

	params1 := mux.Vars(r)
	rollno, _ := strconv.Atoi(params1["rollno"])
	params2 := mux.Vars(r)
	id, _ := strconv.Atoi(params2["id"])

	err := h.d.CheckExist(rollno, id)
	if err != nil {
		fmt.Fprint(w, "student has not been enrolled")
		return
	}
	fmt.Fprint(w, "student has been enrolled")
}

func (h studentHandler) GetNames(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	rollno, _ := strconv.Atoi(params["rollno"])

	resp, err := h.d.GetId(rollno)
	if err != nil {
		fmt.Fprint(w, "some error occured while fetching names")
		return
	}

	fmt.Fprint(w, string(resp))
}
