package main

import (
	"fmt"
)

func checkMagazine(magazine []string, note []string) {
	magazineWords := make(map[string]int)

	// Populate the magazineWords map with the frequency of each word in the magazine
	for _, word := range magazine {
		magazineWords[word]++
	}

	// Check if the note can be constructed using words from the magazine
	for _, word := range note {
		if count, found := magazineWords[word]; found && count > 0 {
			// If the word is present in the magazine, decrement its count
			magazineWords[word]--
		} else {
			// If the word is not present or the count is 0, Print "No" and return
			fmt.Println("No")
			return
		}
	}

	// If all words in the note are present in the magzine, print "Yes"
	fmt.Println("Yes")
}

func main() {
	// Example usage
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}

	checkMagazine(magazine, note)
}
