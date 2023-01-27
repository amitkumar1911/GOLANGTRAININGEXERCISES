package routing

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetPersonDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Failed to connect with database")
	}

	rs := mock.NewRows([]string{"name", "age", "phone"}).AddRow("Amit", 21, "918309172")
	mock.ExpectQuery("SELECT *").WillReturnRows(rs)

	tests := []struct {
		name     string
		database *sql.DB
		want     []Person
		wantErr  error
	}{
		{name: "Success", database: db, want: []Person{{Name: "Amit", Age: 21, Phone: "918309172"}}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPersonDetails(tt.database)
			if err != tt.wantErr {
				t.Errorf("getPersonDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPersonDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertToPerson(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Failed to connect with database")
	}

	// mock.NewRows([]string{"name", "age", "phone"}).AddRow("aman", 21, "23456")
	mockUpdateOutput := sqlmock.NewResult(1, 1)
	mock.ExpectExec("INSERT").WillReturnResult(mockUpdateOutput)


	_,err1:= db.Exec("INSERT INTO person VALUES('aman',21,'23456'")

	tests := []struct {
		name     string
		database *sql.DB
		want     []Person
		wantErr  error
	}{
		{
			name:     "success",
			database: db,
			want:     []Person{{"aman", 21, "23456"}},
			wantErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := insertToPerson(tt.database)
			if err != tt.wantErr {
				t.Errorf("didn't expected error but got one")
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}

}
func Test_rootHandler(t *testing.T) {

	tests := []struct {
		name   string
		input1 *httptest.ResponseRecorder
		input2 *http.Request
		want   string
	}{
		{
			name:   "success case 1",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/ping", nil),
			want:   "pong",
		},
		{
			name:   "success case 2",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/person", nil),
			want:   string(`[{"Name":"amit","Age":21,"Phone":"123"}]`),
		},

		{
			name:   "success case 2",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/hello", nil),
			want:   "invalid",
		},
	}

	a := apiHandler{"mysql", "root:Amit@19sql@tcp(localhost:3306)/recordings"}

	db, _ := sql.Open(a.db, a.connString)

	m := mydb{db}
	for _, tt := range tests {

		m.rootHandler(tt.input1, tt.input2)
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
