package main

import (
	"fmt"
	"math"
)

type nodeftype func(x float64) float64

type fntype func(x, y int) int

func add(a, b int) int {
	return a * b
}

func pow(a, b int) int {
	m := float64(a)
	n := float64(b)
	return int(math.Pow(m, n))
}

func main() {
	var a fntype = add
	var b fntype = pow
	fmt.Printf("%v\n%T\n", a(2, 3), a)
	fmt.Printf("%v\n%T\n", b(2, 3), b)
	var c nodeftype = math.Abs
	fmt.Println(c(3.14))
}

// type Operator func(x float64) float64

// // Map applies op to each element of a.
// func Map(op Operator, a []float64) []float64 {
// 	res := make([]float64, len(a))
// 	for i, x := range a {
// 		res[i] = op(x)
// 	}
// 	return res
// }

// func main() {
// 	op := math.Abs
// 	a := []float64{1, -2}
// 	b := Map(op, a)
// 	fmt.Println(b) // [1 2]

// 	c := Map(func(x float64) float64 { return 10 * x }, b)
// 	fmt.Println(c) // [10, 20]
// }
