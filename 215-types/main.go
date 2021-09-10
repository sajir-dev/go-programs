package main

import (
	"fmt"
	"math"
)

func main() {
	// var h1 hello
	// h1 = "hellooo"
	// fmt.Printf("%T", h1)

	x := -81.34
	fmt.Println(Sqrt(x))
}

// type hello string

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := errNegSqrt(x)
		return 0, e
	}

	return math.Sqrt(x), nil
}

type errNegSqrt float64

func (e errNegSqrt) Error() string {
	return fmt.Sprintf("negative numbers does not have square root %v", float64(e))
}
