package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 4")

	// read input
	readFile, err := os.Open("days/day_4/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	wordSearch := [][]string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		wordSearchLine := []string{}
		for _, char := range line {
			wordSearchLine = append(wordSearchLine, string(char))
		}
		wordSearch = append(wordSearch, wordSearchLine)
	}

	readFile.Close()

	xMax := len(wordSearch[0])
	yMax := len(wordSearch)

	fmt.Println(xMax)
	fmt.Println(yMax)

	total := 0
	for y, wordSearchLine := range wordSearch {
		for x, char := range wordSearchLine {
			if char == "X" {
				params := checkParams{
					WordSearch: wordSearch,
					X:          x,
					Y:          y,
					XMax:       xMax,
					YMax:       yMax,
				}
				total += checkAll(params)
			}
		}
	}

}

// checkAll searches for the word XMAS in all directions.
// Must only be called when
func checkAll(params checkParams) int {
	if params.WordSearch[params.Y][params.X] != "X" {
		panic("Can only check for words when starting on 'X'")
	}

	sum := 0
	sum += checkUp(params)
	sum += checkDown(params)
	sum += checkLeft(params)
	sum += checkRight(params)
	sum += checkUpLeftDiag(params)
	sum += checkDownRightDiag(params)
	return sum
}

func checkUp(params checkParams) int {
	// to check up, require 3 rows above this one to exist
	if params.Y-3 < 0 {
		return 0
	}

	if params.WordSearch[params.Y-1][params.X] == "M" &&
		params.WordSearch[params.Y-2][params.X] == "A" &&
		params.WordSearch[params.Y-3][params.X] == "S" {
		return 1
	}

	return 0
}

func checkDown(params checkParams) int {
	// to check down, require 3 rows below this one to exist
	if params.Y+3 < params.YMax {
		return 0
	}

	if params.WordSearch[params.Y+1][params.X] == "M" &&
		params.WordSearch[params.Y+2][params.X] == "A" &&
		params.WordSearch[params.Y+3][params.X] == "S" {
		return 1
	}

	return 0
}

func checkLeft(params checkParams) int {
	return 0
}

func checkRight(params checkParams) int {
	return 0
}

func checkUpLeftDiag(params checkParams) int {
	return 0
}

func checkUpRightDiag(params checkParams) int {
	return 0
}

func checkDownLeftDiag(params checkParams) int {
	return 0
}

func checkDownRightDiag(params checkParams) int {
	return 0
}

type checkParams struct {
	WordSearch [][]string
	X          int
	Y          int
	XMax       int
	YMax       int
}
