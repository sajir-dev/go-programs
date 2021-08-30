package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := takeInput()
	b := takeInput()

	x := (a[2]/a[1] - b[2]/b[1]) / (a[0]/a[1] - b[0]/b[1])
	y := (a[2]/a[0] - b[2]/b[0]) / (a[0]/a[1] - b[0]/b[1])

	fmt.Println(x, y)
}

func takeInput() []float64 {
	arr := make([]float64, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("enter no:")
		scanner.Scan()
		n := scanner.Text()
		num, err := strconv.ParseFloat(n, 64)
		if err != nil && len(n) > 0 {
			fmt.Println("It should be a number")
		}
		if len(n) != 0 {
			arr = append(arr, num)
		} else {
			break
		}
	}
	return arr[:3]
}
