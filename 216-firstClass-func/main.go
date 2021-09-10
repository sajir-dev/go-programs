package main

import (
	"fmt"
)

func abc(x func(y string) string) func(a string) string {
	fmt.Println("inside abc")
	return x
}

func main() {
	s := "hello"
	xyz := func(a string) string {
		fmt.Println("inside xyz")
		return a
	}

	ff := abc(xyz)
	fmt.Println(ff(s))
}
