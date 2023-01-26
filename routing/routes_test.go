package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// func Test_rootHandler(t *testing.T) {
// 	// type args struct {
// 	// 	w http.ResponseWriter
// 	// 	r *http.Request
// 	// }
// 	// tests := []struct {
// 	// 	name string
// 	// 	args args
// 	// }{
// 	// 	// TODO: Add test cases.
// 	// }
// 	// for _, tt := range tests {
// 	// 	t.Run(tt.name, func(t *testing.T) {
// 	// 		rootHandler(tt.args.w, tt.args.r)
// 	// 	})
// 	// }

// expected := "pong"
// req := httptest.NewRequest(http.MethodGet, "/ping", nil)
// w := httptest.NewRecorder()
// 	rootHandler(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	data, err := ioutil.ReadAll(res.Body)
// 	gotOutput := strings.Replace(string(data), "\n", "", -1)
// 	if err != nil {
// 		t.Errorf("Error: %v", err)
// 	}

// 	if !reflect.DeepEqual(gotOutput, expected) {
// 		t.Errorf("got %q, want %q", gotOutput, expected)
// 	}
// }

func Test_rootHandler(t *testing.T) {

	tests := []struct {
		name   string
		input1 *httptest.ResponseRecorder
		input2 *http.Request
		want   string
	}{
		{
			name:   "get/ping",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/ping", nil),
			want:   "pong",
		},
		{
			name:   "get/person",
			input1: httptest.NewRecorder(),
			input2: httptest.NewRequest(http.MethodGet, "/person", nil),
			want:   string(`[{"Name":"amit","Age":21,"Phone":"123"}]`),
		},
	}

	for _, tt := range tests {

		rootHandler(tt.input1, tt.input2)
		res := (tt.input1).Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		gotOutput := strings.Replace(string(data), "\n", "", -1)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		want := tt.want
		if !reflect.DeepEqual(gotOutput, want) {
			t.Errorf("got %q, want %q", gotOutput, want)
		}

	}
}
