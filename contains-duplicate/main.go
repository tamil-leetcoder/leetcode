package main

import "fmt"

// Brute force
// func containsDuplicate(nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i+1;j<len(nums); j++ {
// 			if nums[i] == nums[j]{
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// optimal
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _,num := range(nums) {
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}

func main () {
	nums := []int{1,2,3,1}
	
	fmt.Println(containsDuplicate(nums))
}