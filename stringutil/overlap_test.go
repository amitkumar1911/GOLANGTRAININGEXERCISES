package stringutil

import (
	"fmt"
	"testing"
)

func TestOverlapOne(t *testing.T) {

	input1 := ""
	input2 := ""

	want := ""

	got := CommonSubstring(input1, input2)

	if want != got {

		fmt.Println("test case fails")
	} else {
		fmt.Println("test case passed")
	}

}

func TestOverlapTwo(t *testing.T) {

	input1 := ""
	input2 := ""

	want := ""

	got := CommonSubstring(input1, input2)

	if want != got {

		fmt.Println("test case fails")
	} else {
		fmt.Println("test case passed")
	}

}
func TestOverlapThree(t *testing.T) {

	input1 := ""
	input2 := "abc"

	want := ""

	got := CommonSubstring(input1, input2)

	if want != got {

		fmt.Println("test case fails")
	} else {
		fmt.Println("test case passed")
	}

}

func TestOverlapFour(t *testing.T) {

	input1 := "abc"
	input2 := ""

	want := ""

	got := CommonSubstring(input1, input2)

	if want != got {

		fmt.Println("test case fails")
	} else {
		fmt.Println("test case passed")
	}

}
