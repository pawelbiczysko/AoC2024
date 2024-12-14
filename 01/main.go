package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var col1 []int
	var col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numsFromLine := strings.Split(scanner.Text(), "   ")

		num1, err := strconv.Atoi(numsFromLine[0])
		if err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", numsFromLine, err)
			continue
		}
		col1 = append(col1, num1)

		num2, err := strconv.Atoi(numsFromLine[1])
		if err != nil {
			fmt.Printf("Error converting %s to integer: %v\n", numsFromLine, err)
			continue
		}
		col2 = append(col2, num2)
	}

	slices.Sort(col1)
	slices.Sort(col2)

	calcDifferenceDistance(col1, col2)
	calcMultipliactionDistance(col1, col2)
}

func calcDifferenceDistance(col1 []int, col2 []int) {
	var distance int = 0

	for i := range col1 {
		var diff int
		if col1[i] < col2[i] {
			diff = col2[i] - col1[i]
		} else {
			diff = col1[i] - col2[i]
		}

		distance = distance + diff
	}

	fmt.Println(distance)
}

func calcMultipliactionDistance(col1 []int, col2 []int) {
	var distance int = 0

	for i := range col1 {
		var numCount = count(
			col2,
			func(x int) bool {
				return x == col1[i]
			})

		distance = distance + col1[i]*numCount
	}

	fmt.Println(distance)
}

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
