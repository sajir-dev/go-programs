package main

import "fmt"

func main() {
	fmt.Println(IsBracketBalancedAdvanced("(()"))
}

func isBracketBalanced(str string) bool {
	bs := []rune(str)

	if len(bs)%2 != 0 {
		return false
	}

	for i := 0; i < len(bs)/2; i++ {
		if bs[i] == 40 {
			if bs[len(bs)-1-i] != 41 {
				return false
			}
		}
	}
	return true
}

func IsBracketBalancedAdvanced(str string) bool {
	bs := []rune(str)
	if len(bs)%2 != 0 || bs[0] == 41 || bs[len(bs)-1] == 40 {
		return false
	}

	i, j := 0, 0

	for i = 0; i < len(bs); i++ {
		// j++ till we find a ")"
		if bs[i] != []rune(")")[0] {
			j++
		} else {
			curr := 0
			for curr < j {
				// check if index is going out
				if i+curr > len(bs)-1 || i-curr < 0 {
					return false
				}

				// check if the characters are open and closes respectively from the index point
				if bs[i-curr-1] != []rune("(")[0] || bs[i+curr] != []rune(")")[0] {
					return false
				}
				curr++
			}
			i = i + j - 1
			j = 0
		}
	}
	return true
}

// (())((()))
