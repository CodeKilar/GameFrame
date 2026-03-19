package main

import "fmt"

func quickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	base := arr[length/2]
	left := make([]int, 0, length)
	equal := make([]int, 0, length)
	right := make([]int, 0, length)

	for _, v := range arr {
		if v < base {
			left = append(left, v)
		} else if v == base {
			equal = append(equal, v)
		} else {
			right = append(right, v)
		}
	}

	result := make([]int, 0, length)
	result = append(result, quickSort(left)...)
	result = append(result, equal...)
	result = append(result, quickSort(right)...)
	return result
}

func mock_quickSort() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	result := quickSort(arr)
	fmt.Println(result)
}
