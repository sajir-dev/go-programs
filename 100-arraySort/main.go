package main

import "fmt"

func main() {
	fmt.Println(arrSort([]int{1, 2, 3}, []int{2, 4, 6}))
}

func arrSort(arr1 []int, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	i, j := 0, 0
	arr := make([]int, m+n-2)

	for m+n > i+j {
		if arr1[i] <= arr2[j] {
			arr[i+j] = arr1[i]
			i++
			if i == len(arr1) {
				arr = append(arr, arr2[j:]...)
				break
			}
		} else {
			arr[i+j] = arr2[j]
			j++
			if j == len(arr2) {
				arr = append(arr, arr1[i:]...)
				break
			}
		}
	}

	return arr
}

// 1 3 5
// 2 4 6
