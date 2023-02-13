package main

import (
	"encoding/csv"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadCsv(f io.Reader, c1 chan []string) {

	r := csv.NewReader(f)

	for {

		record, err := r.Read()

		if err == io.EOF {

			break

		}

		temp := hashing(record[3])
		record[3] = strconv.FormatUint(uint64(temp), 10)
		c1 <- record
	}
	close(c1)

}

func hashing(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func MakeMsg(c1 chan []string, c2 chan string) {

	for i := range c1 {
		msg := fmt.Sprintf("%s%s%s%s", i[0], i[1], i[2], i[3])
		msgSignature := fmt.Sprintf("%-*s", 100, msg)
		c2 <- msgSignature
	}
	close(c2)
}

func Write(c2 chan string, c3 chan string) {

	for i := range c2 {
		temp := hashing(i)
		c3 <- strconv.FormatUint(uint64(temp), 10)

	}
	close(c3)
}

func main() {

	f, err := os.Open("person.csv")
	if err != nil {

		log.Fatalf("%v", err)
	}
	c1 := make(chan []string)
	c2 := make(chan string)
	c3 := make(chan string)
	go ReadCsv(f, c1)
	go MakeMsg(c1, c2)
	go Write(c2, c3)

	csvFile, err := os.Create("textfile.csv")
	csvWriter := csv.NewWriter(csvFile)

	for i := range c3 {
		csvWriter.Write([]string{i})
	}
	csvWriter.Flush()
	csvFile.Close()

}
