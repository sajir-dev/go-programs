package main

import "fmt"

// # 1
// func main() {
// 	c := make(chan int)
// 	go func() {
// 		for {
// 			fmt.Println(<-c)
// 		}
// 	}()
// 	for i := 0; i < 10; i++ {
// 		c <- i
// 	}
// }

// # 2
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

// 1. Go-routine syncing using channels and time.Sleep
// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c := make(chan int)
// 	go odd(c)
// 	go even(c)
// 	time.Sleep(time.Millisecond)
// }

// func odd(c chan int) chan int {
// 	for i:=0;i<100;i++ {
// 		c <- i
// 		if i%2 == 1 {
// 			fmt.Println(i)
// 		}
// 	}
// 	return c
// }

// func even(c chan int) {
// 	for i:=0;i<100;i++ {
// 		<- c
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

// 2. Go-routine syncing using waitgroups and sync packages
// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	c := make(chan int)

// 	go OddNumPrint(c, &wg)
// 	go EvenNumPrint(c, &wg)

// 	wg.Wait()
// }

// func OddNumPrint(c chan int, wg *sync.WaitGroup){
// 	for i:=1; i<100; i++ {
// 		c <- i
// 		if i%2==1{
// 			fmt.Println(i)
// 		}
// 	}
// 	wg.Done()
// }

// func EvenNumPrint(c chan int, wg *sync.WaitGroup){
// 	for i:=1; i<100; i++ {
// 		<- c
// 		if i%2==0 {
// 			fmt.Println(i)
// 		}

// 	}
// 	wg.Done()
// }

// 3. Go routine syncing using errgroup pkg
// import (
// 	"context"
// 	"fmt"
// 	"golang.org/x/sync/errgroup"
// )

// func main() {
// 	g, _ := errgroup.WithContext(context.Background())

// 	g.Go(OddNumPrint())
// 	g.Go(EvenNumPrint())
// 	err := g.Wait()
// 	if err != nil {
// 		fmt.Println("some error occured")
// 	}

// }

// func OddNumPrint() func() error {
// 	return func() error {
// 		for i := 1; i <= 100; i++ {
// 			if i %2 == 1 {
// 				fmt.Println(i)
// 			}
// 		}
// 		return nil
// 	}
// }

// func EvenNumPrint() func() error {
// 	return func() error {
// 		for i := 1; i <= 100; i++ {
// 			if i%2 == 0 {
// 				fmt.Println(i)
// 			}

// 		}
// 		return nil
// 	}
// }

// 4. Go routine syncing using errgroup pkgs
// import (
// 	"context"
// 	"fmt"
// 	"golang.org/x/sync/errgroup"
// )

// func main() {
// 	g, _ := errgroup.WithContext(context.Background())
// 	c := make(chan int)

// 	g.Go(OddNumPrint(c))
// 	g.Go(EvenNumPrint(c))
// 	err := g.Wait()
// 	if err != nil {
// 		fmt.Println("some error occured")
// 	}

// }

// func OddNumPrint(c chan int) func() error {
// 	return func() error {
// 		for i := 1; i <= 100; i++ {
// 			c <- i
// 			if i %2 == 1 {
// 				fmt.Println(i)
// 			}
// 		}
// 		return nil
// 	}
// }

// func EvenNumPrint(c chan int) func() error {
// 	return func() error {
// 		for i := 1; i <= 100; i++ {
// 			<-c
// 			if i%2 == 0 {
// 				fmt.Println(i)
// 			}

// 		}
// 		close(c)
// 		return nil
// 	}
// }
