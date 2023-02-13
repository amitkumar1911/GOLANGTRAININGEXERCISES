package enrollment

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInsert(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MockenrollmentStore
		mockcalls func(*MockenrollmentStore)
		wantErr   error
	}{
		{
			name: "successfull insert operation",
			args: NewMockenrollmentStore(ctrl),
			mockcalls: func(m *MockenrollmentStore) {
				m.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
			wantErr: nil,
		},
		{
			name: "unsuccessfull insert operation",
			args: NewMockenrollmentStore(ctrl),
			mockcalls: func(m *MockenrollmentStore) {
				m.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(errors.New("failed to insert data"))
			},
			wantErr: errors.New("failed to insert data"),
		},
	}

	for _, tt := range tests {
		tt.mockcalls(tt.args)
		enSvc := NewEnrollmentService(tt.args)
		got := enSvc.Insert(4, 5)

		if !reflect.DeepEqual(got, tt.wantErr) {
			t.Errorf("got %q, want %q", got, tt.wantErr)
		}
	}
}

func TestFindRollById(t *testing.T) {

	ctrl := gomock.NewController(t)

	tests := []struct {
		name      string
		args      *MockenrollmentStore
		mockcalls func(*MockenrollmentStore)
		wantValue []int
		wantErr   error
	}{
		{
			name: "successfull case",
			args: NewMockenrollmentStore(ctrl),
			mockcalls: func(m *MockenrollmentStore) {
				m.EXPECT().FindIdByRoll(gomock.Any()).Return([]int{}, nil)
			},
			wantValue: []int{},
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		tt.mockcalls(tt.args)
		enSvc := NewEnrollmentService(tt.args)
		gotValue, gotErr := enSvc.FindIdByRoll(3)

		if !reflect.DeepEqual(gotErr, tt.wantErr) {
			t.Errorf("got %q, want %q", gotErr, tt.wantErr)
		}
		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("got %q, want %q", gotValue, tt.wantErr)
		}
	}
}
