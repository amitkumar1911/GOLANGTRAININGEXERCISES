package nthfibonacci

import (
	"reflect"
	"testing"
)

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

func TestFibonacciNegative(t *testing.T) {

	s := Slc{}

	got := s.Fibonacci(-4)

	want := []int{}

	// fmt.Println(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestFibonacciFour(t *testing.T) {

	s := Slc{}

	got := s.Fibonacci(4)

	want := []int{0, 1, 1, 2}

	// fmt.Println(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestFibonacciFive(t *testing.T) {

	s := Slc{}

	got := s.Fibonacci(5)

	want := []int{0, 1, 1, 2, 3}

	// fmt.Println(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestFibonacciThree(t *testing.T) {

	s := Slc{}

	got := s.Fibonacci(3)

	// got := value(3)

	want := []int{0, 1, 1}

	// fmt.Println(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
