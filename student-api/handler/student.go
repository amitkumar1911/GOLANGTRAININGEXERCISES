package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/student-api/models"
)

type datastore interface {
	Insert(models.Student) error
	GetAll() ([]byte, error)
	Get(rollno int) ([]byte, error)
	Update(rno int, s models.Student) error
}

type studenthandler struct {
	d datastore
}

func NewHandler(db datastore) studenthandler {
	return studenthandler{db}
}

func (h studenthandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := h.d.GetAll()

	if err != nil {
		fmt.Fprint(w, "failed to get all records")
		return
	}
	fmt.Fprint(w, string(resp))
}

func (h studenthandler) Insert(w http.ResponseWriter, r *http.Request) {

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
	err2 := h.d.Insert(s)
	if err2 != nil {
		fmt.Fprint(w, "cannot insert record")
		return
	}
	fmt.Fprint(w, "inserted data successfully")

}

func (h studenthandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value, _ := strconv.Atoi(params["id"])
	resp, err := h.d.Get(value)
	if err != nil {
		fmt.Fprint(w, "failed to get record based on roll no")
		return
	}
	fmt.Fprint(w, string(resp))
}

func (h studenthandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	value, _ := strconv.Atoi(params["id"])
	var s models.Student
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "cannot convert to slice of bytes")
		return
	}
	err1 := json.Unmarshal(bytes, &s)
	if err1 != nil {
		fmt.Fprint(w, "some error occured while parsing json data")
		return
	}

	err2 := h.d.Update(value, s)
	if err2 != nil {
		fmt.Fprint(w, "failed to update record")
		return
	}
	fmt.Fprint(w, "record updated successfully")
}
