package main

import "fmt"

var hello = []string{"Hello", "สวัสดี ครับ", "你好", "😀"}

func main() {
	for _, h := range hello {
		fmt.Printf("%s, world!\n", h)
	}
}
