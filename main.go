package main

import (
	"fmt"
	"net/http"
	"strings"
)

var counted map[rune]int = map[rune]int{'A': 1,
	'E': 1,
	'I': 1,
	'L': 1,
	'N': 1,
	'O': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'U': 1,
	'D': 2,
	'G': 2,
	'M': 2,
	'B': 3,
	'C': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'J': 8,
	'Q': 8,
	'K': 10,
	'W': 10,
	'X': 10,
	'Y': 10,
	'Z': 10}

func count(str string) (value int) {
	value = 0

	for _, v := range strings.ToUpper(str) {
		if _, ok := counted[v]; ok {
			value += counted[v]
		}
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%d", count(r.URL.Path[1:]))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4141", nil)
}
