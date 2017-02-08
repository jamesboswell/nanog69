package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Lang is programming languages
type Lang struct {
	name string
	url  string
}

var langs = []Lang{
	Lang{name: "Go", url: "http://golang.org"},
	Lang{name: "Python", url: "http://python.org"},
	Lang{name: "Perl", url: "https://www.perl.org/"}}

func main() {
	start := time.Now()
	results := make(chan string)
	n := 0
	for _, l := range langs {
		n++
		go count(l.name, l.url, results)
	}

	for i := 0; i < n; i++ {
		fmt.Print(<-results)
	}

	fmt.Println("Total time: ", time.Since(start))
}

func count(name string, url string, results chan string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	results <- fmt.Sprintf("%s\t%d bytes \t[%s]\n", name, n, time.Since(start))
}
