package main

import "fmt"

func main() {
	// whoRuns(2, 10)
	showvalue([]int{2, 3, 5, 8}, 10)
}

func whoRuns(a, b int) {
	if a+b > 10 || a-b < -10 {
		fmt.Println("You are here")
	}
}

func showvalue(sl []int, m int) {
	if m > len(sl)-1 || m < 0 && sl[m] != 3 {
		fmt.Println("Beyond index")
	}
}
