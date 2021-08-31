package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.co.in")
	fmt.Println(res, err)
}
