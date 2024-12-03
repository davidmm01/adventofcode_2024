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
	fmt.Println("Day 1")
	// part 1
	readFile, err := os.Open("days/day_1/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	leftList := []int{}
	rightList := []int{}

	for fileScanner.Scan() {
		temp := strings.Split(fileScanner.Text(), "   ")
		leftListElement, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		rightListElement, err := strconv.Atoi(temp[1])
		if err != nil {
			// ... handle error
			panic(err)
		}

		leftList = insertInOrder(leftListElement, leftList)
		rightList = insertInOrder(rightListElement, rightList)
	}
	readFile.Close()

	if len(leftList) != len(rightList) {
		panic("messed up something, lists should be same length")
	}

	totalDistance := 0

	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance *= -1
		}
		totalDistance += distance
	}
	fmt.Println("Part 1:", totalDistance)

	// part 2
	// just reuse the ordered lists we built up, doesn't change anything
	rightOccuranceMap := buildOccuranceMap(rightList)
	similarityScoreTotal := 0
	for _, element := range leftList {
		similarity := 0
		occurances, ok := rightOccuranceMap[element]
		if ok {
			similarity = element * occurances
		}
		similarityScoreTotal += similarity
	}

	fmt.Println("Part 2:", similarityScoreTotal)
}

func buildOccuranceMap(source []int) map[int]int {
	occuranceMap := make(map[int]int)
	for _, element := range source {
		_, ok := occuranceMap[element]
		if ok {
			occuranceMap[element]++
		} else {
			occuranceMap[element] = 1
		}
	}
	return occuranceMap
}

func insertInOrder(new int, slice []int) []int {
	for index, current := range slice {
		if new <= current {
			return slices.Insert(slice, index, new)
		}
	}
	// either list was empty, or this is the biggest, chuck it at the end
	return append(slice, new)
}
