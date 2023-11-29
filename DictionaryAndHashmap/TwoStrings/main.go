package main

import "fmt"

func twoStrings(s1 string, s2 string) string {
	// Create a map to store the presence of characters in s1
	charactersInS1 := make(map[rune]bool)

	// Populate the map with characters from s1
	for _, char := range s1 {
		charactersInS1[char] = true
	}

	// Check if any character from s2 is present in the map
	for _, char := range s2 {
		if charactersInS1[char] {
			return "YES"
		}
	}

	// If no common substring is found
	return "NO"
}

func main() {
	// Example usage
	fmt.Println(twoStrings("hello", "world")) // Should print "YES"
	fmt.Println(twoStrings("hi", "world"))    // Should print "NO"
}
