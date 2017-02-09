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
	name     string
	url      string
	loadtime time.Duration
}

// URLs with mock load times for Go Playground
var langs = []Lang{
	Lang{name: "Go", url: "http://golang.org", loadtime: parse("86.900303ms")},
	Lang{name: "Python", url: "http://python.org", loadtime: parse("207.981306ms")},
	Lang{name: "Perl", url: "https://www.perl.org/", loadtime: parse("376.4593ms")}}

func main() {
	start := time.Now()
	results := make(chan string)
	n := 0
	// START OMIT
	// inside main()
	for _, l := range langs {
		n++
		// 'go' keyword starts mockhttpGet in a goroutine
		// goroutines are 'green' threads that operate concurrently
		go mockhttpGet(l.name, l.url, l.loadtime, results)
	}
	// END OMIT
	for i := 0; i < n; i++ {
		fmt.Print(<-results)
	}
	fmt.Println("Total time: ", time.Since(start))
}

func mockhttpGet(name string, url string, loadtime time.Duration, results chan string) {
	start := time.Now()
	// sleep to mock a http.Get
	// playground security doesn't allow TCP ops
	time.Sleep(loadtime)
	results <- fmt.Sprintf("%s\t[%s]\n", name, time.Since(start))
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

func parse(t string) time.Duration {
	dur, _ := time.ParseDuration(t)
	return dur
}
