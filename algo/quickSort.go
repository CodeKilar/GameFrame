package main

import "fmt"

// region 堆排序

// region 快速排序
func quickSort(arr []int) []int {
	/*
		1、选择基准元素
		2、通过遍历，将数组分成两部分：小于基准元素的放在左边，大于基准元素的放在右边
		3、递归地对左右两部分进行排序
	*/
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

func mock_sort() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	result := quickSort(arr)
	fmt.Println(result)
}
