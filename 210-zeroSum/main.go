package main

import "fmt"

func subarraySum(nums []int, k int) int {
	cache := make(map[int]int)

	res := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			res += 1
		}

		if v, ok := cache[sum-k]; ok {
			res += v
		}

		if v, ok := cache[sum]; !ok {
			cache[sum] = 1
		} else {
			cache[sum] = v + 1
		}
	}
	return res
}

func main() {
	a := []int{2, -3, 1, 4, -5, 6, -1}
	k := 1
	fmt.Println(subarraySum(a, k))

}
