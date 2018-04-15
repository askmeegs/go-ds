package main

import "fmt"

func main() {
	t1 := []int{1, 2, 2, 3, 4}
	t2 := []int{3, 4, 5, 12, 20, 100, 1000}
	t3 := []int{-5, -5, -5, 0, 1, 1, 4, 5, 6, 6, 6, 6, 7, 7}
	tests := [][]int{t1, t2, t3}

	for _, t := range tests {
		fmt.Println(t)
		result := deleteDuplicates(t)
		fmt.Println(result)
		fmt.Println("\n\n")
	}
}

// Given a sorted array, return an array with the duplicates deleted
// MUST be O(n) time
// OPTIMIZATION - do it in O(1) space by shifting items in place O(1) is the index keeping track of

// delete in place
func deleteDuplicates(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	j := 0 // j is the index of the next unique elt

	/*
	  1 2 2 3 4

	  j = 0
	  i =0, j=0 --> 1 2 2 3 4
	  i = 1, j=1--> 1 2 2 3 4
	  i=2, j=1 -->  1 2 2 3 4
	  i=3, j=2 -->  1 2 3 3 4
	  i=4, j=3 -->  1 2 3 4 4 --> j=4

	  return A[:4] --> return 1 2 3 4

	*/

	for i := 0; i < len(A)-1; i++ {
		if A[i] != A[i+1] {
			A[j] = A[i]
			j++
		}
	}

	A[j] = A[len(A)-1]
	j++

	fmt.Println(j)
	return A[:j]
}
