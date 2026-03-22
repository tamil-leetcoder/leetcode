// 49. Group Anagram
package main

import "fmt"

//  Brute Force
// func groupAnagrams(strs []string) [][]string {
// 	result := [][]string{}
// 	visited := make([]bool, len(strs))

// 	for i:=0; i< len(strs); i++ {
// 		if visited[i] {
// 			continue
// 		}

// 		group := []string{strs[i]}

// 		sI := strings.Split(strs[i], "")
// 		sort.Strings(sI)

// 		for j := i+1; j< len(strs); j++ {
// 			sJ := strings.Split(strs[j], "")
// 			sort.Strings(sJ)

// 			if strings.Join(sI, "") == strings.Join(sJ, "") {
// 				group = append(group, strs[j])
// 				visited[j] = true
// 			}
// 		}

// 		result = append(result, group)
// 	}

// 	return result
// }

// Sorting Key Approach
// func groupAnagrams(strs []string) [][]string {
// 	groups := make(map[string][]string) // Time complexity: O(n*klogk)

// 	for _, word := range strs {
// 		chars := strings.Split(word, "")
// 		sort.Strings(chars)
// 		key := strings.Join(chars, "")
// 		groups[key] = append(groups[key], word)
// 	}

// 	result := [][]string{}
// 	for _, group := range groups {
// 		result = append(result, group)
// 	}

// 	return result
// }

// Character count array approach
func groupAnagrams (strs []string) [][]string{
	groups := make(map[[26]int][]string) // Time Complexity: O(n*k)

	for _, word := range strs {
		var count [26]int

		for _, char := range word {
			count[char-'a']++
		}

		groups[count] = append(groups[count], word)
	}

	result := [][]string{}

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}


func main() {
	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
}