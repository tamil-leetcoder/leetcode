package main

import "fmt"

// Brute Force
// func isAnagram(s, t string) bool {
// 	sArr := strings.Split(s, "")
// 	tArr := strings.Split(t, "")

// 	sort.Strings(sArr)
// 	sort.Strings(tArr)

// 	return strings.Join(sArr, "") == strings.Join(tArr,"") // true or false
// }

// Optimal
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	// increment
	for _, char := range s {
		count[char]++
	}

	//  Decrement
	for _, char := range t {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "east"
	t := "seat"
	result := isAnagram(s,t)

	fmt.Println(result)
}