package main

import "fmt"

func makeAnagram(a string, b string) int32 {
	// Create arrays to store the count of each character in the string
	countA := make([]int, 26)
	countB := make([]int, 26)

	// Count the occurrences of each character in string a
	for _, char := range a {
		countA[char-'a']++
	}

	// Count the occurrences of each character in string b
	for _, char := range b {
		countB[char-'a']++
	}

	// Calculate the total number of deletions required to make the string anagrams
	deletions := int32(0)
	for i := 0; i < 26; i++ {
		deletions += int32(abs(countA[i] - countB[i]))
	}

	return deletions
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example usage
	result := makeAnagram("cde", "dcf")
	fmt.Println(result) // Should print 2
}
