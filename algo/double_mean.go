package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRedPacketAmount(TotalAmount int, TotalNum int) []int64 {

	remainAmount := int64(TotalAmount)
	remainNum := int64(TotalNum)
	result := make([]int64, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < TotalNum-1; i++ {
		amount := remainAmount * 2 / remainNum
		randAmount := rand.Int63n(int64(amount-1)) + 1
		remainNum -= 1
		remainAmount -= randAmount
		result = append(result, randAmount)
	}
	result = append(result, int64(remainAmount))
	return result
}

func mock_binary_mean() {
	fmt.Println(GetRedPacketAmount(1000, 10))
}
