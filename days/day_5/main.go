package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5")

	// read input
	readFile, err := os.Open("days/day_5/input.txt")
	// readFile, err := os.Open("days/day_5/sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	pageOrder := make(map[int][]int)

	inputs := [][]int{}

	pageOrderComplete := false

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if !pageOrderComplete && line == "" {
			pageOrderComplete = true

			continue
		}

		if !pageOrderComplete {
			numsAsStr := strings.Split(line, "|")
			beforeNum, err := strconv.Atoi(numsAsStr[0])
			if err != nil {
				fmt.Println(pageOrder)
				panic("oops 1")
			}
			afterNum, err := strconv.Atoi(numsAsStr[1])
			if err != nil {
				panic("oops 2")
			}

			pageOrder[beforeNum] = append(pageOrder[beforeNum], afterNum)
		} else {
			numsAsStr := strings.Split(line, ",")

			ints := []int{}
			for _, numStr := range numsAsStr {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic("could turn str to int")
				}
				ints = append(ints, num)
			}

			inputs = append(inputs, ints)
		}
	}

	readFile.Close()

	part1Sum := 0

	incorrectlyOrderedInputs := [][]int{}

	for _, input := range inputs {
		ok := true

		// for each number in the input, we will check if the number before it is breaking any rules. if it is not, then continue
		// onto the next number. By this logic, we start comparing element at index 1 with index 0
		// NOTE1!!! This way of checking was not enough for part 2, couldnt just check the current number with previous, had
		// to go through all previous numbers. This is probably a bug in my solution but they made the part 1 easy enough that you wouldn't hit this case until part 2
		for iCurrent := 1; iCurrent < len(input); iCurrent++ {
			iPrevious := iCurrent - 1
			current := input[iCurrent]
			previous := input[iPrevious]

			for _, candidate := range pageOrder[current] {
				if candidate == previous {
					ok = false
				}
			}

			if !ok {
				incorrectlyOrderedInputs = append(incorrectlyOrderedInputs, input)
				break
			}
		}

		// if we didnt break any rules, add the middle number to the sum
		if ok {
			middleIndex := (len(input) - 1) / 2
			part1Sum += input[middleIndex]
		}
	}

	fmt.Println("Part 1:", part1Sum)

	part2Sum := 0
	for _, input := range incorrectlyOrderedInputs {
		allOk := false

		for !allOk {
			for iCurrent := 1; iCurrent < len(input); iCurrent++ {
				// as per NOTE1 above, extra loop here
				for iPrevious := 0; iPrevious < iCurrent; iPrevious++ {
					current := input[iCurrent]
					previous := input[iPrevious]

					for _, candidate := range pageOrder[current] {
						if candidate == previous {
							input[iCurrent], input[iPrevious] = input[iPrevious], input[iCurrent]

							iCurrent = 1
						}
					}
				}

				allOk = true
			}
		}

		// fmt.Println(input)
		middleIndex := (len(input) - 1) / 2
		part2Sum += input[middleIndex]
	}
	fmt.Println("Part 2:", part2Sum)
}
