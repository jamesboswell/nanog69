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
	for _, l := range langs {
		mockhttpGet(l.name, l.url, l.loadtime)
	}
	fmt.Println("Total time: ", time.Since(start))
}

func mockhttpGet(name string, url string, loadtime time.Duration) {
	start := time.Now()
	// sleep to mock a http.Get, play.golang.org disallows external calls
	time.Sleep(loadtime)
	fmt.Printf("%s\t[%s]\n", name, time.Since(start))
}

func count(name, url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	fmt.Printf("%s\t%d bytes \t[%s]\n", name, n, time.Since(start))
}

func parse(t string) time.Duration {
	dur, _ := time.ParseDuration(t)
	return dur
}
