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

	// part1
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
				total += checkAllP1(params)
			}
		}
	}
	fmt.Println("Part 1:", total)

	// part2
	total = 0
	for y, wordSearchLine := range wordSearch {
		for x, char := range wordSearchLine {
			if char == "A" {
				params := checkParams{
					WordSearch: wordSearch,
					X:          x,
					Y:          y,
					XMax:       xMax,
					YMax:       yMax,
				}
				total += checkAllP2(params)
			}
		}
	}
	fmt.Println("Part 2:", total)
}

func checkAllP2(params checkParams) int {
	if params.WordSearch[params.Y][params.X] != "A" {
		panic("Can only check for words when starting on 'A'")
	}

	if params.Y-1 < 0 || params.Y+1 >= params.YMax || params.X-1 < 0 || params.X+1 >= params.XMax {
		return 0
	}

	topLeftToBottomRight := (params.WordSearch[params.Y-1][params.X-1] == "M" && params.WordSearch[params.Y+1][params.X+1] == "S") ||
		(params.WordSearch[params.Y-1][params.X-1] == "S" && params.WordSearch[params.Y+1][params.X+1] == "M")

	bottomLeftToTopRight := (params.WordSearch[params.Y+1][params.X-1] == "M" && params.WordSearch[params.Y-1][params.X+1] == "S") ||
		(params.WordSearch[params.Y+1][params.X-1] == "S" && params.WordSearch[params.Y-1][params.X+1] == "M")

	if topLeftToBottomRight && bottomLeftToTopRight {
		return 1
	}

	return 0
}

type checkParams struct {
	WordSearch [][]string
	X          int
	Y          int
	XMax       int
	YMax       int
}

// checkAllP1 searches for the word XMAS in all directions.
// Must only be called when
func checkAllP1(params checkParams) int {
	if params.WordSearch[params.Y][params.X] != "X" {
		panic("Can only check for words when starting on 'X'")
	}

	sum := 0
	sum += checkUpP1(params)
	sum += checkDownP2(params)
	sum += checkLeftP1(params)
	sum += checkRightP1(params)
	sum += checkUpLeftDiagP1(params)
	sum += checkUpRightDiagP1(params)
	sum += checkDownLeftDiagP1(params)
	sum += checkDownRightDiagP1(params)

	return sum
}

func checkUpP1(params checkParams) int {
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

func checkDownP2(params checkParams) int {
	// to check down, require 3 rows below this one to exist
	if params.Y+3 >= params.YMax {
		return 0
	}

	if params.WordSearch[params.Y+1][params.X] == "M" &&
		params.WordSearch[params.Y+2][params.X] == "A" &&
		params.WordSearch[params.Y+3][params.X] == "S" {
		return 1
	}

	return 0
}

func checkLeftP1(params checkParams) int {
	// to check left, require 3 columns to the left to exist
	if params.X-3 < 0 {
		return 0
	}

	if params.WordSearch[params.Y][params.X-1] == "M" &&
		params.WordSearch[params.Y][params.X-2] == "A" &&
		params.WordSearch[params.Y][params.X-3] == "S" {
		return 1
	}

	return 0
}

func checkRightP1(params checkParams) int {
	// to check right, require 3 columns to the left to exist
	if params.X+3 >= params.XMax {
		return 0
	}

	if params.WordSearch[params.Y][params.X+1] == "M" &&
		params.WordSearch[params.Y][params.X+2] == "A" &&
		params.WordSearch[params.Y][params.X+3] == "S" {
		return 1
	}

	return 0
}

func checkUpLeftDiagP1(params checkParams) int {
	// to check up left diag, require 3 up and 3 left
	if params.Y-3 < 0 || params.X-3 < 0 {
		return 0
	}

	if params.WordSearch[params.Y-1][params.X-1] == "M" &&
		params.WordSearch[params.Y-2][params.X-2] == "A" &&
		params.WordSearch[params.Y-3][params.X-3] == "S" {
		return 1
	}

	return 0
}

func checkUpRightDiagP1(params checkParams) int {
	if params.Y-3 < 0 || params.X+3 >= params.XMax {
		return 0
	}

	if params.WordSearch[params.Y-1][params.X+1] == "M" &&
		params.WordSearch[params.Y-2][params.X+2] == "A" &&
		params.WordSearch[params.Y-3][params.X+3] == "S" {
		return 1
	}

	return 0
}

func checkDownLeftDiagP1(params checkParams) int {
	if params.Y+3 >= params.YMax || params.X-3 < 0 {
		return 0
	}

	if params.WordSearch[params.Y+1][params.X-1] == "M" &&
		params.WordSearch[params.Y+2][params.X-2] == "A" &&
		params.WordSearch[params.Y+3][params.X-3] == "S" {
		return 1
	}

	return 0
}

func checkDownRightDiagP1(params checkParams) int {
	if params.Y+3 >= params.YMax || params.X+3 >= params.XMax {
		return 0
	}

	if params.WordSearch[params.Y+1][params.X+1] == "M" &&
		params.WordSearch[params.Y+2][params.X+2] == "A" &&
		params.WordSearch[params.Y+3][params.X+3] == "S" {
		return 1
	}

	return 0
}
