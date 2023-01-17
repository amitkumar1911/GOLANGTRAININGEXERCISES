package studentutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type students struct {
// 	students []student
// }

type student struct {
	Name   string   //`json:"name"`
	Age    int      // `json:"age"`
	Rollno string   //`json:"rollno"`
	Phone  []string //`json:"phone"`
}

func ParsingJson() {
	// for reading data.json
	jsonFile, openerror := os.Open("data.json")

	if openerror != nil {
		fmt.Println(openerror)
	}
	fmt.Println("successfully opened data.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// fmt.Println(byteValue)

	// var s students

	// error := json.Unmarshal(byteValue, &s)

	// if error != nil {
	// 	fmt.Println(s)
	// }

	// var x map[string]interface{}

	var x []student

	error := json.Unmarshal([]byte(byteValue), &x)

	fmt.Println(x, error)

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

	primary, _ := json.Marshal(p)
	_ = ioutil.WriteFile("primary.json", primary, 0644)

	secondary, werror := json.Marshal(s)
	werror = ioutil.WriteFile("secondary.json", secondary, 0644)

	if werror != nil {
		fmt.Println("cannot write to file")
	}

	// file, _ := json.Marshal(x[0])

	// _ = ioutil.WriteFile("p.json", file, 0644)

	// fmt.Println(len(x))

}
