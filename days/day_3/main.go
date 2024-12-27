package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"regexp"
)

func main() {
	fmt.Println("Day 3")

	digitPairRegexStr := `[\d]{1,3}`
	mulRegexStr := fmt.Sprintf(`mul\(%s,%s\)`, digitPairRegexStr, digitPairRegexStr)
	doRegexStr := `do\(\)`
	dontRegexStr := `don't\(\)`

	mulRegex := regexp.MustCompile(mulRegexStr)
	digitPairRegex := regexp.MustCompile(digitPairRegexStr)
	doRegex := regexp.MustCompile(doRegexStr)
	dontRegex := regexp.MustCompile(dontRegexStr)
	regexForAllInstructions := regexp.MustCompile(fmt.Sprintf(`(%s)|(%s)|(%s)`, mulRegexStr, doRegex, dontRegex))

	// read input
	readFile, err := os.Open("days/day_3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		matches := mulRegex.FindAll([]byte(line), -1)

		for _, match := range matches {
			fmt.Println(string(match))
			digits := digitPairRegex.FindAll([]byte(match), 2)
			digit1, _ := strconv.Atoi(string(digits[0]))
			digit2, _ := strconv.Atoi(string(digits[1]))
			sum += digit1 * digit2
		}

	}
	readFile.Close()

	fmt.Println("Part 1:", sum)

	readFile, err = os.Open("days/day_3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner = bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum = 0
	enabled := true

	for fileScanner.Scan() {
		line := fileScanner.Text()

		instructions := regexForAllInstructions.FindAll([]byte(line), -1)
		for _, instruction := range instructions {
			dont := dontRegex.Match(instruction)
			if dont {
				enabled = false
				continue
			}

			do := doRegex.Match(instruction)
			if do {
				enabled = true
				continue
			}

			mul := mulRegex.Match(instruction)
			if mul {
				if enabled {
					digits := digitPairRegex.FindAll(instruction, 2)
					digit1, _ := strconv.Atoi(string(digits[0]))
					digit2, _ := strconv.Atoi(string(digits[1]))
					sum += digit1 * digit2
				}
				continue
			}

			panic("blowing up:" + string(instruction))
		}
	}
	readFile.Close()

	fmt.Println("Part 2:", sum)
}
