package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestFilterByRoll(t *testing.T) {

	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		input1      *httptest.ResponseRecorder
		input2      *http.Request
		want        string
	}{
		{
			name: "success case",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"name", "rollno", "age"}).AddRow("s", 8, 9).AddRow("q", 8, 10)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/rollno/8", nil),
			want:   string(`[{"Name":"s","Rollno":8,"Age":9},{"Name":"q","Rollno":8,"Age":10}]`),
		},
		{
			name: "error case",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT *").WillReturnError(errors.New("cannot perform the query operation"))
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/student/roll/8", nil),
			want:   "something went wrong",
		},
		{
			name: "invalid case with 4 fields",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"rollno", "age"}).AddRow("s", 8)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/rollno/8", nil),
			want:   "cannot process rows",
		},
		{
			name: "nil io reader",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("UPDATE").WillReturnResult(mockUpdatedOutput)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPut, "/student/roll/2", nil),
			want:   "some error occured",
		},
		{
			name: "valid case",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("UPDATE").WillReturnResult(mockUpdatedOutput)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPut, "/student/roll/2", strings.NewReader(`[{"Name":"f","Rollno":3,"Age":4}]`)),
			want:   "data is updated successfully",
		},
		{
			name: "success post request 2",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("UPDATE").WillReturnResult(mockUpdatedOutput)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPut, "/students/roll/2", errReader(0)),
			want:   "test error",
		},

		{
			name: "success post request 2",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE").WillReturnError(errors.New("cannot update data"))
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPut, "/students/roll/2", strings.NewReader(`[{"Name":"s","Rollno":1,"Age":2}]`)),
			want:   "cannot update data",
		},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatal(err)
		}
		m := mydb{db}

		tt.mockClosure(mock)

		m.filterByRoll(tt.input1, tt.input2)
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

func TestProcessStudents(t *testing.T) {

	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		input1      *httptest.ResponseRecorder
		input2      *http.Request
		want        string
	}{
		{
			name: "success post request 1",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT").WillReturnResult(mockUpdatedOutput)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", nil),
			want:   "some error occured",
		},

		{
			name: "success post request 2",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT").WillReturnResult(mockUpdatedOutput)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`[{"Name":"efg","Rollno":3,"Age":4}]`)),
			want:   "data is entered successfully",
		},
		{
			name: "success post request 2",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT").WillReturnResult(mockUpdatedOutput)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", errReader(0)),
			want:   "test error",
		},
		{
			name: "success case",
			mockClosure: func(mock sqlmock.Sqlmock) {

				rs := mock.NewRows([]string{"name", "rollno", "age"}).AddRow("s", 8, 9)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs)

			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/students", nil),
			want:   string(`[{"Name":"s","Rollno":8,"Age":9}]`),
		},
		{
			name: "error case",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT *").WillReturnError(errors.New("cannot perform the query operation"))
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/students", nil),
			want:   "something went wrong",
		},
		{
			name: "invalid case with 4 fields",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"rollno", "age"}).AddRow("s", 8)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs)
			},
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/students", nil),
			want:   "cannot process rows",
		},
	}
	for _, tt := range tests {

		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatal(err)
		}
		m := mydb{db}

		tt.mockClosure(mock)

		m.processStudent(tt.input1, tt.input2)
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
