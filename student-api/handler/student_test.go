package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	models "github.com/student-api/models"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mockdatastore
		input1    *httptest.ResponseRecorder
		input2    *http.Request
		mockcalls func(m *Mockdatastore)
		want      string
	}{
		{
			name:   "unsuccessfull select operation",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/students", nil),
			mockcalls: func(m *Mockdatastore) {
				m.EXPECT().GetAll().Return([]byte{}, errors.New("failed to get all records")).AnyTimes()
			},
			want: "failed to get all records",
		},
		{
			name:   "successfull select operation",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/students", nil),
			mockcalls: func(m *Mockdatastore) {
				m.EXPECT().GetAll().Return([]byte(`[{"Name":"x","Rollno":2,"Age":3}]`), nil).AnyTimes()
			},
			want: string([]byte(`[{"Name":"x","Rollno":2,"Age":3}]`)),
		},
	}

	for _, tt := range tests {
		tt.mockcalls(tt.args)
		studenthandler := NewHandler(tt.args)
		studenthandler.GetAll(tt.input1, tt.input2)
		res := tt.input1.Result()
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		got := strings.Replace(string(data), "\n", "", -1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %q, want %q", got, tt.want)
		}

	}

}

type errReader int

func (e errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestInsert(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mockdatastore
		input1    *httptest.ResponseRecorder
		input2    *http.Request
		mockcalls func(m *Mockdatastore)
		want      string
	}{
		{
			name:   "ioutil.readAll error",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", errReader(0)),
			mockcalls: func(m *Mockdatastore) {
				//nothing to do
			},
			want: "cannot convert to slice of bytes",
		},
		{
			name:   "json.Unmarshal error",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", nil),
			mockcalls: func(m *Mockdatastore) {
				//nothing to do
			},
			want: "some error occured while parsing the json data",
		},
		{
			name:   "unsuccessfull insertion",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Name":"x","Rollno":1,"Age":2}`)),
			mockcalls: func(m *Mockdatastore) {

				m.EXPECT().Insert(models.Student{Name: "x", Rollno: 1, Age: 2}).Return(errors.New("cannot insert record")).AnyTimes()

			},
			want: "cannot insert record",
		},
		{
			name:   "successfull insertion",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Name":"x","Rollno":1,"Age":2}`)),
			mockcalls: func(m *Mockdatastore) {

				m.EXPECT().Insert(models.Student{Name: "x", Rollno: 1, Age: 2}).Return(nil).AnyTimes()

			},
			want: "inserted data successfully",
		},
	}
	for _, tt := range tests {
		tt.mockcalls(tt.args)
		studenthandler := NewHandler(tt.args)
		studenthandler.Insert(tt.input1, tt.input2)
		res := tt.input1.Result()
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		got := strings.Replace(string(data), "\n", "", -1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %q, want %q", got, tt.want)
		}

	}

}

func TestGet(t *testing.T) {

	vars := map[string]string{
		"id": "8",
	}

	ctrl := gomock.NewController(t)
	tests := []struct {
		name      string
		args      *Mockdatastore
		input1    *httptest.ResponseRecorder
		input2    *http.Request
		mockcalls func(m *Mockdatastore)
		want      string
	}{
		{
			name:   "unsuccessfull get operation",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/students/roll/8", nil), vars),
			mockcalls: func(m *Mockdatastore) {
				m.EXPECT().Get(gomock.Any()).Return([]byte{}, errors.New("error occured")).AnyTimes()
			},
			want: "failed to get record based on roll no",
		},
		{
			name:   "successfull get operation",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/students/roll/8", nil), vars),
			mockcalls: func(m *Mockdatastore) {
				m.EXPECT().Get(gomock.Any()).Return([]byte(`[{"Name":"x","Rollno":8,"Age":3}]`), nil).AnyTimes()
			},
			want: string([]byte(`[{"Name":"x","Rollno":8,"Age":3}]`)),
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		studenthandler := NewHandler(tt.args)
		studenthandler.Get(tt.input1, tt.input2)
		res := tt.input1.Result()
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		got := strings.Replace(string(data), "\n", "", -1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %q, want %q", got, tt.want)
		}

	}
}

func TestUpdate(t *testing.T) {

	vars := map[string]string{"id": "8"}
	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mockdatastore
		input1    *httptest.ResponseRecorder
		input2    *http.Request
		mockcalls func(m *Mockdatastore)
		want      string
	}{
		{
			name:   "ioutil.readAll error",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/students/roll/8", errReader(0)), vars),
			mockcalls: func(m *Mockdatastore) {
				//nothing to do
			},
			want: "cannot convert to slice of bytes",
		},
		{
			name:   "json.Unmarshal error",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/students/roll/8", nil), vars),
			mockcalls: func(m *Mockdatastore) {
				//nothing to do
			},
			want: "some error occured while parsing json data",
		},
		{
			name:   "unsuccessfull update opration",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/students/roll/8", strings.NewReader(`{"Name":"x","Age":2}`)), vars),
			mockcalls: func(m *Mockdatastore) {

				m.EXPECT().Update(gomock.Any(), models.Student{Name: "x", Age: 2}).Return(errors.New("failed to update record")).AnyTimes()

			},
			want: "failed to update record",
		},
		{
			name:   "successfull update operation",
			args:   NewMockdatastore(ctrl),
			input1: httptest.NewRecorder(),
			input2: mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/students/roll/8", strings.NewReader(`{"Name":"x","Age":2}`)), vars),
			mockcalls: func(m *Mockdatastore) {

				m.EXPECT().Update(gomock.Any(), models.Student{Name: "x", Age: 2}).Return(nil).AnyTimes()

			},
			want: "record updated successfully",
		},
	}
	for _, tt := range tests {
		tt.mockcalls(tt.args)
		studenthandler := NewHandler(tt.args)
		studenthandler.Update(tt.input1, tt.input2)
		res := tt.input1.Result()
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		got := strings.Replace(string(data), "\n", "", -1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %q, want %q", got, tt.want)
		}

	}
}
