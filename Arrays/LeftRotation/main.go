package main

import "fmt"

func rotLeft(a []int32, d int32) []int32 {
	n := int32(len(a))

	// Calculate the effective number of rotations by taking the modulo of d with the array length
	rotations := d % n

	// Create a new slice to store the rotated array
	rotatedArray := make([]int32, n)

	// Perform the left rotation
	for i := int32(0); i < n; i++ {
		// Calculate the new index after rotation
		newIndex := (i - rotations + n) % n
		rotatedArray[newIndex] = a[i]
	}

	return rotatedArray
}

func main() {
	// Sample Input
	d := int32(4)
	arr := []int32{1, 2, 3, 4, 5}

	// Function call
	result := rotLeft(arr, d)

	// Print the result as a single line of space-separated integers
	for _, num := range result {
		fmt.Printf("%d ", num)
	}

}
