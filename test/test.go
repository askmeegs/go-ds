// testing tool based on codejam ".in" files
package test

import (
	"bufio"
	"fmt"
	"os"
)

// a solver is a function that generates an answer to a programming question
type Solver func(string) (string, error)

// f = test case file
func SolveTestfile(f string, s Solver) error {
	output := []string{}
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result, err := s(scanner.Text())
		if err != nil {
			return fmt.Errorf("Test case [%s] produced err: %v", scanner.Text(), err)
		}
		output = append(output, result)
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
		line := fmt.Sprintf("Case #%d: %s\n", i, o)
		_, err := f.WriteString(line)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote %d test outputs to %s\n", len(output), outfile)
	return nil
}
