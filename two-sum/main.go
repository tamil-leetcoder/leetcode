package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int);

	for i, num := range nums {
		complement := target - num

		if j, found := seen[complement]; found {
			return []int{i,j}
		}

		seen[num] = i
	}

	return nil
}

func main() {
	nums := []int{2,7,11,19}
	target := 9

	indexes := twoSum(nums, target)
	
	fmt.Println(indexes)
}