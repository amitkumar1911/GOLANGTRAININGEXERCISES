package subject

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"strings"
	"testing"

	models "github.com/GOLANGTRAININGEXERCISES/student-api/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

type errReader int

func (e errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestCreate(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MocksubjectService
		w         *httptest.ResponseRecorder
		r         *http.Request
		mockcalls func(*MocksubjectService)
		want      string
	}{
		{
			name: "readAll error",
			args: NewMocksubjectService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/subjects", errReader(0)),
			mockcalls: func(m *MocksubjectService) {
				//nothing to do
			},
			want: "cannot convert to slice of bytes",
		},
		{
			name: "json error",
			args: NewMocksubjectService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", nil),
			mockcalls: func(m *MocksubjectService) {
				//nothing to do
			},
			want: "some error occured while parsing the json data",
		},
		{
			name: "failed insert operation",
			args: NewMocksubjectService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Name":"x","Id":2}`)),
			mockcalls: func(m *MocksubjectService) {
				m.EXPECT().CreateSubject(models.Subject{Name: "x", Id: 2}).Return(errors.New("cannot insert record"))
			},
			want: "cannot insert record",
		},
		{
			name: "successfull insert operation",
			args: NewMocksubjectService(ctrl),
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(`{"Name":"y","Id":3}`)),
			mockcalls: func(m *MocksubjectService) {
				m.EXPECT().CreateSubject(models.Subject{Name: "y", Id: 3}).Return(nil)
			},
			want: "inserted data successfully",
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		subjectHandler := NewSubjectHandler(tt.args)
		subjectHandler.Create(tt.w, tt.r)
		resp := tt.w.Result()
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		got := strings.Replace(string(data), "\n", "", -1)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %s, want %s", got, tt.want)
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
		args      *MocksubjectService
		w         *httptest.ResponseRecorder
		r         *http.Request
		mockcalls func(*MocksubjectService)
		want      string
	}{
		{
			name: "failed get operation",
			args: NewMocksubjectService(ctrl),
			w:    httptest.NewRecorder(),
			r:    mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/subjects", nil), vars),
			mockcalls: func(m *MocksubjectService) {
				m.EXPECT().GetSubject(8).Return([]byte{}, errors.New("failed to get all records"))
			},
			want: "failed to get all records",
		},
		{
			name: "successfull get operation",
			args: NewMocksubjectService(ctrl),
			w:    httptest.NewRecorder(),
			r:    mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/subjects", nil), vars),
			mockcalls: func(m *MocksubjectService) {
				m.EXPECT().GetSubject(8).Return([]byte(`[{"Name":"x", "Id":2}]`), nil)
			},
			want: string([]byte(`[{"Name":"x", "Id":2}]`)),
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		subjectHandler := NewSubjectHandler(tt.args)
		subjectHandler.Get(tt.w, tt.r)
		resp := tt.w.Result()
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		got := strings.Replace(string(data), "\n", "", -1)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf(" got %s, want %s", got, tt.want)
		}

	}
}
