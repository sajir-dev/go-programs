package main

import (
	"fmt"
	"net"
)

func main() {
	// url := "https://google.com"
	// res, err := http.Get(url)
	// fmt.Println(res, err)

	iprecords, _ := net.LookupIP("facebook.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
}
