package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []int{4, 1, 3, 2}
	result := quicksort(a)
	fmt.Println(result)
}

// quicksort
func quicksort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := 0
	right := len(nums) - 1
	pivot := random(len(nums))

	// swap rightmost with pivot
	nums[pivot], nums[right] = nums[right], nums[pivot]

	// pivot is now a[right] and it won't move during our loop...

	// everything smaller than pivot goes to the left
	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	// put the pivot back right where "left" left off. ha.
	nums[left], nums[right] = nums[right], nums[left]

	// recurse!
	quicksort(nums[:left])
	quicksort(nums[left+1:])

	return nums
}

func random(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}
