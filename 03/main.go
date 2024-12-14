package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	byteInput, err := os.ReadFile("test_input_2.txt")
	if err != nil {
		fmt.Print(err)
	}

	strInput := string(byteInput)

	// findAndMultiply(strInput)

	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	// matches := re.FindAllString(strInput, -1)
	// for _, v := range matches {
	// 	fmt.Println(v)
	// }
	strInput = re.ReplaceAllString(strInput, "")
	fmt.Println(strInput)

	findAndMultiply(strInput)
}

func findAndMultiply(strInput string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllString(strInput, -1)

	var sum int = 0

	for _, match := range matches {
		re := regexp.MustCompile(`mul\((.*)\)`)
		args := re.FindStringSubmatch(match)[1]
		nums := strings.Split(args, ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		sum = sum + num1*num2
	}

	fmt.Println(sum)
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
