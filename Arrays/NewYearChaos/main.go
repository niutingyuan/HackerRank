package main

import (
	"fmt"
)

func minimumBribes(q []int32) {
	bribes := int32(0)
	p1 := int32(1)
	p2 := int32(2)
	p3 := int32(3)
	for i := 0; i < len(q); i++ {
		if q[i] == p1 {
			// 1 2 3 originally
			p1 = p2
			p2 = p3
			p3++
			// 2 3 4 after the above assignments
		} else if q[i] == p2 {
			// 2 1 3 originally
			bribes++
			p2 = p3
			p3++
			// 2 3 4 after the above assignments
		} else if q[i] == p3 {
			// 3 1 2 originally
			bribes += 2
			p3++
			// 1 2 4 after the above assignments
		} else {
			fmt.Println("Too chaotic")
			return
		}
	}
	fmt.Println(bribes)
}

func main() {
	// Example usage:
	q1 := []int32{1, 2, 3, 5, 4, 6, 7, 8}
	minimumBribes(q1) // Output: 1

	q2 := []int32{4, 1, 2, 3}
	minimumBribes(q2) // Output: Too chaotic
}
