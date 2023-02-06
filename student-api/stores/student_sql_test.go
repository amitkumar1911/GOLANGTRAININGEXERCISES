package stores

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/student-api/models"
)

func TestGetAll(t *testing.T) {

	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		wantValue   []byte
		wantErr     error
	}{
		{
			name: "unsuccessfull select operation",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT *").WillReturnError(errors.New("failed to select rows"))
			},
			wantValue: nil,
			wantErr:   errors.New("failed to select rows"),
		},
		{
			name: "scan rows errror",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"rollno", "age"}).AddRow(2, 3)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs).WillReturnError(errors.New("some error occured while scanning rows"))
			},
			wantValue: nil,
			wantErr:   errors.New("some error occured while scanning rows"),
		},
		{
			name: "succesfull select operation",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"name", "rollno", "age"}).AddRow("x", 2, 3)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs).WillReturnError(nil)
			},
			wantValue: []byte(`[{"Name":"x","Rollno":2,"Age":3}]`),
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := SqlDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		tt.mockClosure(mock)

		gotValue, gotErr := d.GetAll()

		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantValue)
		}
		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)

		}
	}

}
func TestInsert(t *testing.T) {

	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		input       models.Student
		wantErr     error
	}{
		{
			name: "successfull insertion",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT").WithArgs("x", 1, 2).WillReturnResult(mockUpdatedOutput)
			},
			input:   models.Student{Name: "x", Rollno: 1, Age: 2},
			wantErr: nil,
		},
		{
			name: "unsuccessfull insertion",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT").WithArgs("", 0, 0).WillReturnError(errors.New("failed to insert data successfully"))
			},
			input:   models.Student{},
			wantErr: errors.New("failed to insert data successfully"),
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := SqlDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// defer db.Close()
		tt.mockClosure(mock)

		gotErr := d.Insert(tt.input)

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}

	}

}

func TestGet(t *testing.T) {

	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		input       int
		wantValue   []byte
		wantErr     error
	}{
		{
			name: "unsuccessfull query operation",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT *").WithArgs(215).WillReturnError(errors.New("failed to select rows"))
			},
			input:     215,
			wantValue: nil,
			wantErr:   errors.New("failed to select rows"),
		},
		{
			name: "scan rows errror",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"rollno", "age"}).AddRow(2, 3)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs).WillReturnError(errors.New("some error occured while scanning rows"))
			},
			input:     215,
			wantValue: nil,
			wantErr:   errors.New("some error occured while scanning rows"),
		},
		{
			name: "succesfull select operation",
			mockClosure: func(mock sqlmock.Sqlmock) {
				rs := mock.NewRows([]string{"name", "rollno", "age"}).AddRow("y", 216, 316)
				mock.ExpectQuery("SELECT *").WithArgs(216).WillReturnRows(rs).WillReturnError(nil)
			},
			input:     216,
			wantValue: []byte(`[{"Name":"y","Rollno":216,"Age":316}]`),
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := SqlDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		tt.mockClosure(mock)

		gotValue, gotErr := d.Get(tt.input)

		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantValue)
		}
		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)

		}
	}

}

func TestUpdate(t *testing.T) {

	type args struct {
		rollno int
		s      models.Student
	}

	db, mock, err := sqlmock.New()
	d := SqlDb{db}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	tests := []struct {
		Name        string
		args        args
		mockClosure func(mock sqlmock.Sqlmock)
		wantErr     error
	}{
		{
			Name: "error case",
			args: args{1, models.Student{Name: "x", Age: 2}},
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE").WithArgs("x", 2, 1).WillReturnError(errors.New("unsuccessfull update operation"))
			},
			wantErr: errors.New("unsuccessfull update operation"),
		},
		{
			Name: "success case",
			args: args{1, models.Student{Name: "x", Age: 2}},
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE").WithArgs("x", 2, 1).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {

		tt.mockClosure(mock)

		gotErr := d.Update(tt.args.rollno, tt.args.s)

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}

	}

}
