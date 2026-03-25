package main

import "fmt"

/*
1、建堆
2、向下拓展
*/

func buildHeap(nums []int) {
	// 从最后一个非叶子节点，从下往上堆排序
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

func heapify(nums []int, n, i int) {
	for {
		largest := i
		left := 2*i + 1  //左节点
		right := 2*i + 2 // 右节点

		if left < n && nums[left] > nums[largest] {
			largest = left
		}

		if right < n && nums[right] > nums[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}

func heapSort(nums []int) {
	buildHeap(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		// 堆大小-1，从新的堆顶开始调整
		heapify(nums, i, 0)

	}

}

func mock_heap() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	heapSort(arr)
	fmt.Println(arr)
}
