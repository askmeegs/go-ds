package main

import (
	"strconv"
	"strings"

	"github.com/m-okeefe/go-ds/test"
)

const LIMIT = 100000

// https://code.google.com/codejam/contest/6254486/dashboard#s=p0

func main() {
	test.SolveTestfile("A-large-practice.in", SheepBrute)
}

/*
Problem:  "Counting Sheep"

Take in a number N
Then say N, 2N, 3N, 4N ....
And look at the digits in each of those numbers
Keep track of which digits (in total) seen at least once
Once you've seen all ten digits 0-9 at least once, terminate

GIVEN: N
OUTPUT: the last number you say before terminating (reaching goal)
        OR "INSOMNIA" if you will count forever

0 1 2 3 4 5 6 7 8 9


*/

/*
Patterns
-- has nothing to do w/ N being prime.
-- or with Odd / Even (1234567890 is even and fine.)

0 -> insomnia
1 -> 1,2,3,4... (10)
2 -> 2,4,6,8,10,12,14 --> fine
3 -> 3,6,9,12,15,18,21,24,27,30 --> fine.
*/

func SheepBrute(in string) string {
	N, err := strconv.Atoi(in)
	if err != nil {
		return err.Error()
	}

	found := []bool{} //array / zero valued to false
	for i := 0; i < 10; i++ {
		found = append(found, false)
	}

	for i := 1; i < LIMIT; i++ {
		cI := N * i
		cS := strconv.Itoa(cI)
		digits := strings.Split(cS, "")
		update(digits, found)
		if done(found) {
			return cS
		}
	}
	return "INSOMNIA"
}

func update(digits []string, found []bool) []bool {
	for _, i := range digits {
		d, _ := strconv.Atoi(i)
		found[d] = true
	}
	return found
}

func done(found []bool) bool {
	for _, b := range found {
		if b == false {
			return false
		}
	}
	return true
}
