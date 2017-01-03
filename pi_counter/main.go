package main

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

// TODO - Document
func getPiDigit(i int) (int8, error) {

	pi_digits := [105]int8{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 2, 7, 9, 5, 0, 2, 8, 8, 4, 1, 9, 7, 1, 6, 9, 3, 9, 9, 3, 7, 5, 1, 0, 5, 8, 2, 0, 9, 7, 4, 9, 4, 4, 5, 9, 2, 3, 0, 7, 8, 1, 6, 4, 0, 6, 2, 8, 6, 2, 0, 8, 9, 9, 8, 6, 2, 8, 0, 3, 4, 8, 2, 5, 3, 4, 2, 1, 1, 7, 0, 6, 7, 9, 8, 2, 1, 4}

	if i >= len(pi_digits) {
		return -1, errors.New("I don't have that digit")
	}

	return pi_digits[i], nil
}

// TODO - Document
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

// TODO - Document
func digitHandler(w http.ResponseWriter, r *http.Request) {

	pathFinder := regexp.MustCompile("^/digit/([0-9]+)$")

	match := pathFinder.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.NotFound(w, r)
		return
	}

	index, _ := strconv.Atoi(match[1])

	digit, err := getPiDigit(index)

	if err != nil {
		fmt.Fprintf(w, "?")
	} else {
		fmt.Fprintf(w, "%d", digit)

	}
}

func main() {
	// TODO - Read from config file?
	listenAddr := ":8080"

	// Main Index Page
	http.HandleFunc("/", indexHandler)

	// Digit Handler
	http.HandleFunc("/digit/", digitHandler)

	// Static Files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start Serving files
	fmt.Println("Listening on", listenAddr)
	http.ListenAndServe(listenAddr, nil)
}
