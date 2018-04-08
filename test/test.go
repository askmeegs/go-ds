// testing tool based on codejam ".in" files
package test

import (
	"bufio"
	"fmt"
	"os"
)

// a solver is a function that generates an answer to a programming question
type Solver func(string) string

// f = test case file
func SolveTestfile(f string, s Solver) error {
	output := []string{}
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //trash 1st line (# of test cases)
	i := 1
	for scanner.Scan() {
		// fmt.Printf("Solving case #%d\n", i)
		result := s(scanner.Text())
		output = append(output, result)
		i = i + 1
	}
	return writeTestResults(output, fmt.Sprintf("OUT_%s", f))
}

func writeTestResults(output []string, outfile string) error {
	f, err := os.Create(outfile)
	defer f.Close()
	if err != nil {
		return err
	}
	for i, o := range output {
		line := fmt.Sprintf("Case #%d: %s\n", i+1, o)
		_, err := f.WriteString(line)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote %d test outputs to %s\n", len(output), outfile)
	return nil
}
