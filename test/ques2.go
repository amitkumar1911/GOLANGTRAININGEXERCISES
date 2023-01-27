package main

import (
	"fmt"
	"net/http"
)

func printRequest(count int, name string) string {

	res := ""

	if count == 1 {

		res = res + "Greetings" + name

	} else {

		res = res + "welcome back" + name

	}
	return res

}

func keeptrack() func(http.ResponseWriter, *http.Request) {

	count := 0

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(count)

		count++

		if count == 1 {

			name := r.URL.Path
			ans := printRequest(count, name)
			w.Write([]byte(ans))

		} else {

			name := r.URL.Path
			ans := printRequest(count, name)
			w.Write([]byte(ans))

		}
	}
}

func main() {

	f := keeptrack()

	err := http.ListenAndServe(":8000", http.HandlerFunc(f))

	if err != nil {
		fmt.Println("something went wrong")
	}
}
