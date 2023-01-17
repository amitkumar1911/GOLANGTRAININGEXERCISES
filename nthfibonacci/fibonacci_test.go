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
func TestFibonacciTwo(t *testing.T) {

	s := Slc{}

	got := s.Fibonacci(2)

	// value := Fibonacci()

	want := []int{0, 1}

	// fmt.Println(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

// func TestFibonacciNegative(t *testing.T) {

// 	s := Slc{}

// 	got := s.Fibonacci(-4)

// 	want := []int{}

// 	// fmt.Println(got)

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}

// }

// func TestFibonacciFour(t *testing.T) {

// 	s := Slc{}

// 	got := s.Fibonacci(4)

// 	want := []int{0, 1, 1, 2}

// 	// fmt.Println(got)

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}

// }

// func TestFibonacciFive(t *testing.T) {

// 	s := Slc{}

// 	got := s.Fibonacci(5)

// 	want := []int{0, 1, 1, 2, 3}

// 	// fmt.Println(got)

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}

// }

// func TestFibonacciThree(t *testing.T) {

// 	s := Slc{}

// 	got := s.Fibonacci(3)

// 	// got := value(3)

// 	want := []int{0, 1, 1}

// 	// fmt.Println(got)

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}

// }
