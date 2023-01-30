package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetStudents(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	rs := mock.NewRows([]string{"name", "rollno", "age"}).AddRow("s", 8, 9)
	mock.ExpectQuery("SELECT *").WillReturnRows(rs)

	tests := []struct {
		name   string
		input1 *httptest.ResponseRecorder
		input2 *http.Request
		want   string
	}{
		{name: "success case",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/students", nil),
			want:   string(`[{"Name":"s","Rollno":8,"Age":9}]`),
		},
	}
	m := mydb{db}
	for _, tt := range tests {

		m.getStudents(tt.input1, tt.input2)
		res := (tt.input1).Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		gotOutput := strings.Replace(string(data), "\n", "", -1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if !reflect.DeepEqual(gotOutput, tt.want) {
			t.Errorf("got %q, want %q", gotOutput, tt.want)
		}

	}

}

func TestStudentWithGivenRoll(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatal(err)
	}

	rs := mock.NewRows([]string{"name", "rollno", "age"}).AddRow("s", 8, 9).AddRow("q", 8, 10)
	mock.ExpectQuery("SELECT *").WillReturnRows(rs)

	tests := []struct {
		name   string
		input1 *httptest.ResponseRecorder
		input2 *http.Request
		want   string
	}{
		{name: "success case",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/rollno?rollno=8", nil),
			want:   string(`[{"Name":"s","Rollno":8,"Age":9},{"Name":"q","Rollno":8,"Age":10}]`),
		},
	}
	m := mydb{db}
	for _, tt := range tests {

		m.getStudents(tt.input1, tt.input2)
		res := (tt.input1).Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		gotOutput := strings.Replace(string(data), "\n", "", -1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if !reflect.DeepEqual(gotOutput, tt.want) {
			t.Errorf("got %q, want %q", gotOutput, tt.want)
		}

	}
}
