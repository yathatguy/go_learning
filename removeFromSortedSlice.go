package main

import (
	"fmt"
)

func main() {
	length := removeDuplicates([]int{0, 1, 1, 2, 2, 2})
	fmt.Println(length)
}

func removeDuplicates(nums []int) int {
	var (
		cnt int
	)
	for i := 0; i < len(nums)-1; i++ {
		for nums[i] == nums[i+1] {
			for j := i + 1; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			if i+cnt == len(nums) {
				fmt.Println(nums[:len(nums)-cnt+1])
				return len(nums[:len(nums)-cnt+1])
			}
			cnt += 1
		}
	}
	return len(nums)
}
