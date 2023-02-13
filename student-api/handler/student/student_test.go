package student

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"strings"
	"testing"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

type errReader int

func (e errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestGet(t *testing.T) {

	vars := map[string]string{
		"rollno": "8",
	}
	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MockstudentService
		w         *httptest.ResponseRecorder
		r         *http.Request
		mockcalls func(*MockstudentService)
		want      string
	}{
		{
			name: "unsuccessfull get operation 1",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/students/", nil), vars),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().GetStudent(gomock.Any()).Return([]byte{}, errors.New("failed to get all records")).AnyTimes()
			},
			want: "failed to get all records",
		},

		{
			name: "successfull get operation",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/students/", nil), vars),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().GetStudent(gomock.Any()).Return([]byte(`[{"Name":"x","Rollno":8,"Age":2}]`), nil).AnyTimes()
			},
			want: string([]byte(`[{"Name":"x","Rollno":8,"Age":2}]`)),
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		studentHandler := NewStudentHandler(tt.args)
		studentHandler.GetStudent(tt.w, tt.r)
		resp := tt.w.Result()
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		got := strings.Replace(string(data), "\n", "", -1)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %s, want %s", got, tt.want)
		}

	}

}

func TestCreate(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MockstudentService
		w         *httptest.ResponseRecorder
		r         *http.Request
		mockcalls func(*MockstudentService)
		want      string
	}{
		{
			name: "readAll error",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", errReader(0)),
			mockcalls: func(m *MockstudentService) {
				//nothing to do
			},
			want: "cannot convert to slice of bytes",
		},
		{
			name: "json error",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", nil),
			mockcalls: func(m *MockstudentService) {
				//nothing to do
			},
			want: "some error occured while parsing the json data",
		},
		{
			name: "failed insert operation",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Name":"x","Rollno":1,"Age":2}`)),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().CreateStudent(models.Student{Name: "x", Rollno: 1, Age: 2}).Return(errors.New("cannot insert record"))
			},
			want: "cannot insert record",
		},
		{
			name: "successfull insert operation",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Name":"x","Rollno":1,"Age":2}`)),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().CreateStudent(models.Student{Name: "x", Rollno: 1, Age: 2}).Return(nil)
			},
			want: "inserted data successfully",
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		studentHandler := NewStudentHandler(tt.args)
		studentHandler.CreateStudent(tt.w, tt.r)
		resp := tt.w.Result()
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		got := strings.Replace(string(data), "\n", "", -1)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %s, want %s", got, tt.want)
		}

	}

}

func TestEnrollSubject(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MockstudentService
		w         *httptest.ResponseRecorder
		r         *http.Request
		mockcalls func(*MockstudentService)
		want      string
	}{
		{
			name: "unsuccessfull enrollment",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students/2/subjects/3", nil),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().CheckExist(gomock.Any(), gomock.Any()).Return(errors.New("some error occured")).AnyTimes()
			},
			want: "student has not been enrolled",
		},
		{
			name: "successfull enrollment",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students/2/subjects/3", nil),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().CheckExist(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
			want: "student has been enrolled",
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		studentHandler := NewStudentHandler(tt.args)
		studentHandler.EnrollStudent(tt.w, tt.r)
		resp := tt.w.Result()
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		got := strings.Replace(string(data), "\n", "", -1)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %s, want %s", got, tt.want)
		}

	}

}

func TestGetNames(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MockstudentService
		w         *httptest.ResponseRecorder
		r         *http.Request
		mockcalls func(*MockstudentService)
		want      string
	}{
		{
			name: "failed to fetch names",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/students/2/subjects", nil),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().GetId(gomock.Any()).Return([]byte{}, errors.New("some error occured"))
			},
			want: "some error occured while fetching names",
		},
		{
			name: "names fetched successfully",
			args: NewMockstudentService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/students/2/subjects", nil),
			mockcalls: func(m *MockstudentService) {
				m.EXPECT().GetId(gomock.Any()).Return([]byte(`["Name":"amit"]`), nil)
			},
			want: string([]byte(`["Name":"amit"]`)),
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		studentHandler := NewStudentHandler(tt.args)
		studentHandler.GetNames(tt.w, tt.r)
		resp := tt.w.Result()
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		got := strings.Replace(string(data), "\n", "", -1)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %s, want %s", got, tt.want)
		}

	}
}
