package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(sortString(str))
}

func sortString(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] > bs[j] {
				bs[i], bs[j] = bs[j], bs[i]
			}
		}
	}
	return string(bs)
}
