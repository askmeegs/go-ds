package main

import (
	"fmt"

	"github.com/m-okeefe/go-ds/ds"
)

func main() {
}

func arraySlice() {
	ds.ArrayBasics()
	ds.SliceBasics()
}

func linkedList() {
	L := ds.DoublyLinkedList{}
	L.Insert("a") // a
	L.Insert("b") // b -> a
	L.Insert("c") // c -> b -> a
	L.Insert("d") // d -> c -> b -> a
	// fmt.Println(L.ToString())

	R := L.Reverse()
	// fmt.Println(R.ToString())

	R.Delete("c") //a -> b -> d
	fmt.Println(R.ToString())

	R.Delete("a") // b-> d
	fmt.Println(R.ToString())

	R.Delete("d") // b
	fmt.Println(R.ToString())

	R.Delete("z") // b
	fmt.Println(R.ToString())

	R.Delete("b") // empty
}

func stack() {
	S := ds.Stack{}
	S.Push("fox")
	S.Push("brown")
	S.Push("quick")

	fmt.Println(S.ToString()) // quick brown fox

	S.Pop()                   //brown fox
	S.Pop()                   // fox
	S.Push("lazy")            // lazy fox
	fmt.Println(S.ToString()) // lazy fox
}
