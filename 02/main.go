package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var safeCount int = 0
	var partialSafeCount int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportLevels, _ := parseInts(scanner.Text())

		if isReportSafe(reportLevels) {
			safeCount++
		}

		for i := range reportLevels {
			newSlice := make([]int, len(reportLevels))
			copy(newSlice, reportLevels)
			newSlice = append(newSlice[:i], newSlice[i+1:]...)

			if isReportSafe(newSlice) {
				partialSafeCount++
				break
			}
		}
	}

	fmt.Println(safeCount)
	fmt.Println(partialSafeCount)
}

func isReportSafe(levels []int) bool {
	var raising bool = levels[0] < levels[1]
	for i := range len(levels) - 1 {
		if levels[i] == levels[i+1] {
			return false
		} else if levels[i] < levels[i+1] && !raising {
			return false
		} else if levels[i] > levels[i+1] && raising {
			return false
		} else if levels[i]-levels[i+1] > 3 || levels[i]-levels[i+1] < -3 {
			return false
		}
	}

	return true
}

func parseInts(s string) ([]int, error) {
	var (
		fields = strings.Fields(s)
		ints   = make([]int, len(fields))
		err    error
	)

	for i, f := range fields {
		ints[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}
