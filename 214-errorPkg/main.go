package main

import (
	"fmt"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "errorOne happened"
}

// func (e *errorOne) Error() string {
// 	return "errorOne happened"
// }

func main() {
	e1 := errorOne{}
	e2 := &errorOne{}

	fmt.Println(e1)
	fmt.Println(e2)
}

// NOTES: when a type has function named Error(), it implements error interface and hence it returns the error value attached as the string provided.
// Value receivers can receive both value and pointer types but pointer type can receive only pointers
