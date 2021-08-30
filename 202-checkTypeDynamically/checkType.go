package main

import "fmt"

func main() {
	whoAreYou(123.431)
}

func whoAreYou(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer type")
	case float64:
		fmt.Println("Float type")
	case string:
		fmt.Println("string type")
	default:
		fmt.Printf("%T", v)
	}
}
