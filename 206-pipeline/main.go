package main

import "fmt"

func main() {
	c1 := maker(11, 2, 40)
	c2 := consumer(c1)

	for v := range c2 {
		fmt.Println(v)
	}
}

func maker(nums ...int) chan int {
	c := make(chan int)
	go func() {
		for _, v := range nums {
			c <- v / 2
			fmt.Println()
		}
		close(c)
	}()
	return c
}

func consumer(c chan int) chan int {
	cr := make(chan int)
	go func() {
		for v := range c {
			cr <- v * v
		}
		close(cr)
	}()
	return cr
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	c1 := maker()
// 	c2 := consumer(c1)
// 	for v := range c2 {
// 		fmt.Println(v)
// 	}
// }

// func maker() chan int {
// 	c1 := make(chan int)
// 	go func(){
// 		for i := 0; i< 10; i++ {
// 			c1 <- i*2
// 			fmt.Println("from maker")
// 		}
// 		close(c1)
// 	}()
// 	return c1
// }

// func consumer(c chan int) chan int {
// 	c2 := make(chan int)
// 	go func() {
// 		for v := range c {
// 			c2 <- v * v
// 			fmt.Println("from consumer")
// 		}
// 		close(c2)
// 	}()
// 	return c2
// }
