package subject

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
)

func TestGetSubject(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		mockClosure func(mock sqlmock.Sqlmock)
		wantValue   []byte
		wantErr     error
	}{
		{
			name:  "unsuccessfull select operation",
			input: 2,
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT *").WithArgs(2).WillReturnError(errors.New("failed to select rows"))
			},
			wantValue: nil,
			wantErr:   errors.New("failed to select rows"),
		},
		{
			name:  "scan rows errror",
			input: 3,
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"rollno"}).AddRow(2)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs).WillReturnError(errors.New("some error occured while scanning rows"))
			},
			wantValue: nil,
			wantErr:   errors.New("some error occured while scanning rows"),
		},
		{
			name:  "succesfull select operation",
			input: 2,
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"name", "id"}).AddRow("x", 2)
				mock.ExpectQuery("SELECT *").WithArgs(2).WillReturnRows(rs).WillReturnError(nil)
			},
			wantValue: []byte(`[{"Name":"x","Id":2}]`),
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := SubjectDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		tt.mockClosure(mock)

		gotValue, gotErr := d.GetSubject(tt.input)

		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantValue)
		}
		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)

		}
	}

}

func TestCreateSubject(t *testing.T) {
	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		input       models.Subject
		wantErr     error
	}{
		{
			name: "successfull insertion",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT").WithArgs("x", 1).WillReturnResult(mockUpdatedOutput)
			},
			input:   models.Subject{Name: "x", Id: 1},
			wantErr: nil,
		},
		{
			name: "unsuccessfull insertion",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT").WithArgs("", 0).WillReturnError(errors.New("failed to insert data successfully"))
			},
			input:   models.Subject{},
			wantErr: errors.New("failed to insert data successfully"),
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := SubjectDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// defer db.Close()
		tt.mockClosure(mock)

		gotErr := d.CreateSubject(tt.input)

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}
	}
}

func TestCheckSubjectExist(t *testing.T) {
	tests := []struct {
		name        string
		input       int
		mockClosure func(sqlmock.Sqlmock)
		want        bool
	}{
		{
			name:  "unsuccessfull operation",
			input: 2,
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"Name", "Id"}).AddRow("y", 2)
				mock.ExpectQuery("SELECT COUNT").WithArgs(2).WillReturnRows(rs)
			},
			want: false,
		},

		{
			name:  "successfull operation",
			input: 2,
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"Count"}).AddRow(1)
				mock.ExpectQuery("SELECT COUNT").WithArgs(2).WillReturnRows(rs)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := SubjectDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		tt.mockClosure(mock)
		got := d.CheckSubjectExist(tt.input)

		if got != tt.want {
			t.Errorf("got %q", "different boolean value")
		}

	}
}
