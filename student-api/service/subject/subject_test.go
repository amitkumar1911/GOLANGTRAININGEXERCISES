package subject

import (
	"errors"
	"reflect"
	"testing"

	"github.com/GOLANGTRAININGEXERCISES/student-api/models"
	"github.com/golang/mock/gomock"
)

func TestCreateSubject(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mocksubjectstore
		mockcalls func(*Mocksubjectstore)
		input     models.Subject
		wantErr   error
	}{
		{
			name: "unsuccessfull create operation 1",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				//nothing to do
			},
			input:   models.Subject{Name: "", Id: 2},
			wantErr: errors.New("invalid request"),
		},
		{
			name: "unsuccessfull create operation 1",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().CreateSubject(gomock.Any()).Return(errors.New("cannot create new subject")).AnyTimes()
			},
			input:   models.Subject{Name: "x", Id: 2},
			wantErr: errors.New("cannot create new subject"),
		},
		{
			name: "successfull create operation 1",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().CreateSubject(gomock.Any()).Return(nil).AnyTimes()
			},
			input:   models.Subject{Name: "x", Id: 2},
			wantErr: nil,
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		subSvc := NewSubjectService(tt.args)

		got := subSvc.CreateSubject(tt.input)
		if !reflect.DeepEqual(got, tt.wantErr) {

			t.Errorf("got %q, want %q", got, tt.wantErr)
		}
	}

}

func TestGetSubject(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mocksubjectstore
		mockcalls func(*Mocksubjectstore)
		input     int
		wantValue []byte
		wantErr   error
	}{
		{
			name: "unsuccesfull get operation 1",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				// nothing to do
			},
			input:     0,
			wantValue: []byte{},
			wantErr:   errors.New("invalid request"),
		},
		{
			name: "unsuccessfull get operation 2",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().GetSubject(gomock.Any()).Return([]byte{}, errors.New("failed to get subjects"))
			},
			input:     1,
			wantValue: []byte{},
			wantErr:   errors.New("failed to get subjects"),
		},
		{
			name: "successfull get operation",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().GetSubject(gomock.Any()).Return([]byte(`[{"Name":"maths","Id":2}]`), nil)
			},
			input:     1,
			wantValue: []byte(`[{"Name":"maths","Id":2}]`),
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		tt.mockcalls(tt.args)
		stuSvc := NewSubjectService(tt.args)
		gotValue, gotErr := stuSvc.GetSubject(tt.input)

		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantValue)
		}

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}

	}

}

func TestSubjectExist(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mocksubjectstore
		mockcalls func(*Mocksubjectstore)
		input     int
		want      bool
	}{
		{
			name: "unsuccesfull check 1",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().CheckSubjectExist(gomock.Any()).Return(false).AnyTimes()
			},
			input: 1,
			want:  false,
		},

		{
			name: "succesfull check",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().CheckSubjectExist(gomock.Any()).Return(true).AnyTimes()
			},
			input: 2,
			want:  true,
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		subSvc := NewSubjectService(tt.args)
		got := subSvc.SubjectExist(tt.input)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %q", "different boolean value")
		}
	}
}

func TestFindNamesById(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *Mocksubjectstore
		mockcalls func(*Mocksubjectstore)
		wantValue []byte
		wantErr   error
	}{
		{
			name: "successfull operation",
			args: NewMocksubjectstore(ctrl),
			mockcalls: func(m *Mocksubjectstore) {
				m.EXPECT().FindNamesById(gomock.Any()).Return([]byte{}, nil).AnyTimes()
			},
			wantValue: []byte{},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {

		tt.mockcalls(tt.args)
		subSvc := NewSubjectService(tt.args)
		gotValue, gotErr := subSvc.FindNamesById([]int{2})

		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantValue)
		}
		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)

		}
	}

}
