package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c2 := makeEven()
// 	c1 := makeOdd()
// 	for v := range c1 {
// 		fmt.Println(v)
// 	}
// 	for v := range c2 {
// 		fmt.Println(v)
// 	}
// }

// func makeEven() chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 1; i <= 100; i++ {
// 			if i%2 == 0 {
// 				c <- i
// 			}
// 		}
// 		close(c)
// 	}()
// 	return c
// }func groutine1(p chan int) {

// func makeOdd() chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 1; i <= 100; i++ {
// 			if i%2 != 0 {
// 				c <- i
// 			}
// 		}
// 		close(c)
// 	}()var POOL = 100

// 	return c
// }

// print 1 to 100 using two go-routines

// func odd(ch chan bool, wg *sync.WaitGroup) {
// 	for i := 1; i < 100; i = i + 2 {
// 		<-ch
// 		fmt.Println(i)
// 		time.Sleep(1 * time.Millisecond)
// 		ch <- true
// 	}
// 	wg.Done()
// }

// func even(ch chan bool, wg *sync.WaitGroup) {
// 	for i := 2; i <= 100; i = i + 2 {
// 		<-ch
// 		fmt.Println(i)
// 		time.Sleep(1 * time.Millisecond)
// 		ch <- true
// 	}
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	c := make(chan bool)

// 	wg.Add(2)
// 	go even(c, &wg)
// 	go odd(c, &wg)

// 	c <- true
// 	wg.Wait()
// }

// var POOL = 100

// func groutine1(p chan int) {

// 	for i := 1; i <= POOL; i++ {
// 		p <- i
// 		if i%2 == 1 {
// 			fmt.Println("groutine-1:", i)
// 		}
// 	}
// }

// func groutine2(p chan int) {

// 	for i := 1; i <= POOL; i++ {
// 		<-p
// 		if i%2 == 0 {
// 			fmt.Println("groutine-2:", i)
// 		}
// 	}
// }

// func main() {
// 	msg := make(chan int)

// 	go groutine1(msg)
// 	go groutine2(msg)

// 	time.Sleep(time.Second * 1)

// }

var limit = 100

func odd(c chan int) {
	for i := 1; i <= limit; i++ {
		c <- i
		// fmt.Println(i)
		if i%2 == 1 {
			fmt.Println("Goroutine:1", i)
		}
	}
}

func even(c chan int) {
	for i := 2; i <= limit; i++ {
		<-c
		// fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Goroutine:2", i)
		}
	}
}

func main() {
	c := make(chan int)
	go odd(c)
	go even(c)
	time.Sleep(1 * time.Millisecond)
}
