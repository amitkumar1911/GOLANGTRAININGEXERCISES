package stringutil

import (
	"testing"
)

func TestCommonSubstring(t *testing.T) {

	errorTestCases := []struct {
		name     string
		input1   string
		input2   string
		expected string
	}{
		{name: "bothemptystring", input1: "", input2: "", expected: ""},
		{name: "onestringempty", input1: "", input2: "abc", expected: ""},
		{name: "commonsubstringatstart", input1: "amitkumar", input2: "amit", expected: "amit"},
		{name: "commonsubstringatend", input1: "amitkumar", input2: "kumar", expected: "kumar"},
	}

	for _, value := range errorTestCases {

		got := CommonSubstring(value.input1, value.input2)

		want := value.expected

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}

	}
}
