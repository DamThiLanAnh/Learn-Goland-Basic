package main

import "fmt"

func sumV2(nums ...int) int {
	var result = 0
	for i := 0; i <= len(nums); i++ {
		result += nums[i]
	}
	return result
}

func main() {
	var result = sumV2(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
