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
	for _, l := range langs {
		go count(l.name, l.url, results)
	}

	for c := range results {
		fmt.Print(c)
	}

	fmt.Println("Total time: ", time.Since(start))
}

func count(name string, url string, results *chan string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	results <- fmt.Sprintf("%s\t%d bytes \t[%s]\n", name, n, time.Since(start))
}
