package main

import "fmt"

func main() {
	c := fact(4)
	for v := range c {
		fmt.Println(v)
	}
}

func fact(n int) chan int {
	c := make(chan int)
	go func() {
		res := 1
		for i := n; i > 1; i-- {
			res = res * i
		}
		c <- res
		close(c)
	}()
	return c
}
