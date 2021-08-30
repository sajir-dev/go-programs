package main

import "fmt"

func main() {
	var st string
	fmt.Scanf("%s", &st)
	fmt.Println(charFreq(st))
}

func charFreq(s string) map[string]int {
	freq := make(map[string]int)

	bs := []byte(s)

	for _, v := range bs {
		if freq[string(v)] == 0 {
			freq[string(v)] = 1
		} else {
			freq[string(v)]++
		}
	}
	return freq
}
