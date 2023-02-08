package subject

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
	"github.com/gorilla/mux"
)

type subjectService interface {
	CreateSubject(models.Subject) error
	GetSubject(int) ([]byte, error)
}

type subjectHandler struct {
	d subjectService
}

func NewSubjectHandler(s subjectService) subjectHandler {
	return subjectHandler{s}
}

func (h subjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value, _ := strconv.Atoi(params["id"])
	resp, err := h.d.GetSubject(value)

	if err != nil {
		fmt.Fprint(w, "failed to get all records")
		return
	}
	fmt.Fprint(w, string(resp))
}

func (h subjectHandler) Create(w http.ResponseWriter, r *http.Request) {

	var s models.Subject
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
	err2 := h.d.CreateSubject(s)
	fmt.Println(err2)
	if err2 != nil {
		fmt.Fprint(w, "cannot insert record")
		return
	}
	fmt.Fprint(w, "inserted data successfully")

}
