package main

import (
	"fmt"

	"github.com/m-okeefe/go-ds/ds"
)

func main() {

}

func linkedList() {
	L := ds.DoublyLinkedList{}
	L.Insert("a") // a
	L.Insert("b") // b -> a
	L.Insert("c") // c -> b -> a
	L.Insert("d") // d -> c -> b -> a
	fmt.Println(L.ToString())

	R := L.Reverse()
	fmt.Println(R.ToString())
}

func arraySlice() {
	ds.ArrayBasics()
	ds.SliceBasics()
}
