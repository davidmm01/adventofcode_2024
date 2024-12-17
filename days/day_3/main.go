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

	mulRegex := regexp.MustCompile(`mul\([[\d]{1,3},[\d]{1,3}\)`)
	digitPairRegex := regexp.MustCompile(`[\d]{1,3}`)

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
}
