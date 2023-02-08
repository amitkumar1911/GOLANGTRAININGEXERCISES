package student

import (
	"errors"
	"reflect"
	"testing"

	models "github.com/GOLANGTRAININGEXERCISES/student-api/models"
	"github.com/golang/mock/gomock"
)

func TestCreateStudent(t *testing.T) {

	type args struct {
		stuStr *Mockstudentstore
		subSvc *MocksubjectService
		enrSvc *MockenrollmentService
	}

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      args
		mockcalls func(*Mockstudentstore, *MocksubjectService, *MockenrollmentService)
		input     models.Student
		wantErr   error
	}{
		{
			name: "unsuccessfull get operation 1",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				//nothing to mock
			},
			input:   models.Student{Name: "", Rollno: 2, Age: 3},
			wantErr: errors.New("invalid request"),
		},
		{
			name: "unsuccessfull insert operation 2",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().Create(gomock.Any()).Return(errors.New("failed to insert data")).AnyTimes()
			},
			input:   models.Student{Name: "Rakesh", Rollno: 2, Age: 3},
			wantErr: errors.New("failed to insert data"),
		},

		{
			name: "successfull insert operation",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()
			},
			input:   models.Student{Name: "Rakesh", Rollno: 2, Age: 3},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		tt.mockcalls(tt.args.stuStr, tt.args.subSvc, tt.args.enrSvc)
		stuSvc := NewStudentService(tt.args.stuStr, tt.args.subSvc, tt.args.enrSvc)
		got := stuSvc.CreateStudent(tt.input)

		if !reflect.DeepEqual(got, tt.wantErr) {
			t.Errorf("got %q, want %q", got, tt.wantErr)
		}

	}
}

func TestGetStudent(t *testing.T) {

	ctrl := gomock.NewController(t)

	type args struct {
		stuStr *Mockstudentstore
		subSvc *MocksubjectService
		enrSvc *MockenrollmentService
	}
	tests := []struct {
		name      string
		args      args
		mockcalls func(*Mockstudentstore, *MocksubjectService, *MockenrollmentService)
		input     int
		wantValue []byte
		wantErr   error
	}{
		{
			name: "unsuccessfull get operation 1",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				//nothing to do
			},
			input:     0,
			wantValue: []byte{},
			wantErr:   errors.New("invalid request"),
		},

		{
			name: "unsuccessfull get operation 2",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().Get(gomock.Any()).Return([]byte{}, errors.New("cannot get the data")).AnyTimes()
			},
			input:     2,
			wantValue: []byte{},
			wantErr:   errors.New("cannot get the data"),
		},
		{
			name: "successfull get operation",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().Get(gomock.Any()).Return([]byte(`[{"Name":"x","Rollno":2,"Age":3}]`), nil).AnyTimes()
			},
			input:     2,
			wantValue: []byte(`[{"Name":"x","Rollno":2,"Age":3}]`),
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		tt.mockcalls(tt.args.stuStr, tt.args.subSvc, tt.args.enrSvc)
		stuSvc := NewStudentService(tt.args.stuStr, tt.args.subSvc, tt.args.enrSvc)
		gotValue, gotErr := stuSvc.GetStudent(tt.input)

		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantValue)
		}

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}

	}
}

func TestCheckExist(t *testing.T) {

	ctrl := gomock.NewController(t)

	type args struct {
		stuStr *Mockstudentstore
		subSvc *MocksubjectService
		enrSvc *MockenrollmentService
	}

	type input struct {
		rollno int
		id     int
	}

	tests := []struct {
		name      string
		args      args
		mockcalls func(*Mockstudentstore, *MocksubjectService, *MockenrollmentService)
		input     input
		wantErr   error
	}{
		{
			name: "unsuccessfull check 1",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().StudentExist(gomock.Any()).Return(false).AnyTimes()
			},
			input:   input{2, 3},
			wantErr: errors.New("invalid params"),
		},

		{
			name: "unsuccessfull check 2",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().StudentExist(gomock.Any()).Return(true).AnyTimes()
				mockSubject.EXPECT().SubjectExist(gomock.Any()).Return(false).AnyTimes()
			},
			input:   input{2, 3},
			wantErr: errors.New("invalid params"),
		},
		{
			name: "unsuccessfull check 3",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().StudentExist(gomock.Any()).Return(true).AnyTimes()
				mockSubject.EXPECT().SubjectExist(gomock.Any()).Return(true).AnyTimes()
				mockEnr.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(errors.New("some error occured")).AnyTimes()
			},
			input:   input{2, 3},
			wantErr: errors.New("some error occured"),
		},
		{
			name: "successfull check ",
			args: args{NewMockstudentstore(ctrl), NewMocksubjectService(ctrl), NewMockenrollmentService(ctrl)},
			mockcalls: func(mockStudent *Mockstudentstore, mockSubject *MocksubjectService, mockEnr *MockenrollmentService) {
				mockStudent.EXPECT().StudentExist(gomock.Any()).Return(true).AnyTimes()
				mockSubject.EXPECT().SubjectExist(gomock.Any()).Return(true).AnyTimes()
				mockEnr.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
			input:   input{2, 3},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		tt.mockcalls(tt.args.stuStr, tt.args.subSvc, tt.args.enrSvc)
		stuSvc := NewStudentService(tt.args.stuStr, tt.args.subSvc, tt.args.enrSvc)
		got := stuSvc.CheckExist(tt.input.rollno, tt.input.id)

		if !reflect.DeepEqual(got, tt.wantErr) {
			t.Errorf("got %q, want %q", got, tt.wantErr)
		}

	}
}
