package enrollment

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsert(t *testing.T) {

	type args struct {
		rollno int
		id     int
	}

	tests := []struct {
		name        string
		mockClosure func(sqlmock.Sqlmock)
		args        args
		wantErr     error
	}{
		{
			name: "successfull insertion",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mockUpdatedOutput := sqlmock.NewResult(1, 1)
				mock.ExpectExec("INSERT").WithArgs(2, 3).WillReturnResult(mockUpdatedOutput)
			},
			args:    args{2, 3},
			wantErr: nil,
		},
		{
			name: "unsuccessfull insertion",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT").WithArgs(3, 4).WillReturnError(errors.New("failed to insert data successfully"))
			},
			args:    args{3, 4},
			wantErr: errors.New("failed to insert data successfully"),
		},
	}

	for _, tt := range tests {
		db, mock, err := sqlmock.New()
		d := EnrollDb{db}

		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		// defer db.Close()
		tt.mockClosure(mock)

		gotErr := d.Insert(tt.args.rollno, tt.args.id)

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}
	}
}
