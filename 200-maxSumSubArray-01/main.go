package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}

	fmt.Println(maxSubArray(a))
}

func maxSubArray(nums []int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	fmt.Println(total)
	left := 0
	right := len(nums) - 1
	var curr int
	// result := nums
	result, curr := total, total
	for left < right {
		if nums[left] >= nums[right] {
			curr = curr - nums[right]
			right--
		} else {
			curr = curr - nums[left]
			left++
		}
		if curr > result {
			// result := nums[left:right+1]
			fmt.Println(curr)
			result = curr
		}
	}
	return result
}
