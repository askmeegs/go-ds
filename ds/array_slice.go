package ds

import (
	"fmt"
)

// Arrays = not so common in Go
// they cannot change size!
func ArrayBasics() {
	var A [5]int //go arrays are zero valued fixed size
	A[3] = 6
	fmt.Println(A)

	B := [3]float32{3.14, 5.5, 0.01}
	for i := 0; i < len(B); i++ {
		B[i] = B[i] * 2
	}
	fmt.Println(B)

	// NOTE - this is an array and NOT a slice, because length is specified
	C := [2]string{"hello", "world"}
	D := [...]string{"a", "b", "c"}
	fmt.Println(C)
	fmt.Println(D)
}

// Slices = flexible and more common
func SliceBasics() {

	// Slices are zero valued at NIL
	N := []int{}
	fmt.Println(N)

	// non fixed length
	A := []string{}
	A = append(A, "hi")
	fmt.Println(A)

	// fixed length (make len )
	B := make([]string, 5)
	for i := range B {
		B[i] = fmt.Sprintf("item-%d", i)
	}
	fmt.Println(B)

	//make obj len cap (capacity)
	C := make([]int, 2, 4)
	for i := 0; i < 5; i++ {
		C = append(C, i)
		fmt.Printf("C=%v    Mem=%p     Cap=%d\n", C, C, cap(C))
	}
}
