package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randNum := rand.Intn(100) + 1
	fmt.Println("随机生成的整数：", randNum)
	target := randNum
	i := 1
	j := 100
	for i <= j {
		mid := (i + j) / 2
		if mid == target {
			fmt.Println("找到目标整数：", target)
			break
		} else if mid < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
}
