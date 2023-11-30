package main

import (
	"fmt"
	"strings"
)

func stringConstruction(s string) int32 {
	// Initialize variables to keep track of the cost and the constructed string
	cost := int32(0)
	constructedString := ""

	// Iterate through each character in the input string
	for _, char := range s {

		// Check if the character is already present in the constructed string
		if strings.ContainsRune(constructedString, char) {
			// If yes, append the substring from the constructed string at no cost
			constructedString += string(char)
		} else {
			// If no, append the character to the constructed string at a cost of 1 dollar
			constructedString += string(char)
			cost++
		}
	}

	return cost
}

func main() {
	result := stringConstruction("abcabc")
	fmt.Println(result)
}
