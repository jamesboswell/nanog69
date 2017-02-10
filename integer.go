package main

import "fmt"

func main() {
	number := 16
	fmt.Println(number)
	number = "42" // this will fail and not compile!
	fmt.Println(number)
}
