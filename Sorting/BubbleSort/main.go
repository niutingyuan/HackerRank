package main

import "fmt"

func countSwaps(a []int32) {
	n := len(a)
	var numSwaps int

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				// Swap elements
				a[j], a[j+1] = a[j+1], a[j]
				numSwaps++
			}
		}
	}

	// Print the required output
	fmt.Printf("Array is sorted in %d swaps. \n", numSwaps)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[n-1])
}

func main() {
	// Example usage
	arr := []int32{3, 2, 1}
	countSwaps(arr)
}
