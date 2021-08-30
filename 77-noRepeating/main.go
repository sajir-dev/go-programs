package main

import "fmt"

func main() {
	var st string
	fmt.Scanf("%s", &st)
	fmt.Println(nonRepeatChar(st))
}

func nonRepeatChar(s string) []string {
	bs := []byte(s)
	ss := make([]string, 0)
	m := make(map[byte]int)

	for _, v := range bs {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			ss = append(ss, string(k))
		}
	}
	return ss
}
