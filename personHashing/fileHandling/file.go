package fileHandling

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/GOLANGTRAININGEXERCISES/personHashing/hashPhoneno"
)

func ReadFromCsv(f io.Reader, c1 chan []string) {

	r, _ := csv.NewReader(f).ReadAll()

	for i := 0; i < len(r); i++ {
		temp := hashPhoneno.Hashing(r[i][3])
		r[i][3] = strconv.FormatUint(uint64(temp), 10)
		c1 <- r[i]
	}
	close(c1)
}

func WriteToCsv(w io.Writer, data string) {

	w.Write([]byte(data + "\n"))
}
