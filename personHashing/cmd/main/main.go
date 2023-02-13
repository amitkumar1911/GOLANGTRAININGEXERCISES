package main

import (
	"log"
	"os"

	"github.com/GOLANGTRAININGEXERCISES/personHashing/fileHandling"
	"github.com/GOLANGTRAININGEXERCISES/personHashing/msgPadding"
)

func main() {

	f, err := os.Open("person.csv")
	f1, err := os.Create("data.csv")
	if err != nil {

		log.Fatalf("%v", err)
	}

	defer f1.Close()

	c1 := make(chan []string)
	c2 := make(chan string)
	go fileHandling.ReadFromCsv(f, c1)
	go msgPadding.MakeMsg(c1, c2)

	for value := range c2 {
		fileHandling.WriteToCsv(f1, value)
	}

}
