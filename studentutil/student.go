package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type student struct {
	Name   string
	Age    int
	Rollno string
	Phone  []string
}

func main() {
	// for reading data.json
	jsonFile, openerror := os.Open("data.json")

	if openerror != nil {
		fmt.Println(openerror)
	}
	fmt.Println("successfully opened data.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var x []student

	err := json.Unmarshal([]byte(byteValue), &x)

	fmt.Println(x, err)

	length := len(x)

	var p []student

	var s []student

	for i := 0; i < length; i++ {

		if x[i].Age >= 5 {
			s = append(s, x[i])
		} else {
			p = append(p, x[i])
		}
	}

	primary, perr := json.Marshal(p)
	perr = ioutil.WriteFile("primary.json", primary, 0644)

	if perr != nil {
		fmt.Println("cannot write to file")
	}

	secondary, serr := json.Marshal(s)
	serr = ioutil.WriteFile("secondary.json", secondary, 0644)

	if serr != nil {
		fmt.Println("cannot write to file")
	}

}
