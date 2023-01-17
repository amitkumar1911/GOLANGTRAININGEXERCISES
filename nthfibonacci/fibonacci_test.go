package nthfibonacci

import (
	"reflect"
	"testing"
)

func tableDrivenTest(t *testing.T) {

	s := Slc{}

	errorTestCases := []struct {
		name     string
		input    int
		expected []int
	}{{name: "nequalszero", input: 0, expected: []int{}},
		{name: "nequalstwo", input: 2, expected: []int{0, 1}},
		{name: "nequals1", input: 1, expected: []int{0}},
		{name: "nequalsnegative", input: -2, expected: []int{}},
		{name: "nequals5", input: 5, expected: []int{0, 1, 1, 2, 3}},
		{name: "nequals8", input: 8, expected: []int{0, 1, 1, 2, 3, 5, 8, 13}},
		{name: "nequals4", input: 4, expected: []int{0, 1, 1, 2}}}

	for _, value := range errorTestCases {

		got := s.Fibonacci(value.input)

		want := value.expected

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

}
