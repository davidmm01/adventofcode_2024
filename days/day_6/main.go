package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	OBSTRUCTION = "#"

	DOWN  = "v"
	UP    = "^"
	LEFT  = "<"
	RIGHT = ">"
)

func main() {
	fmt.Println("Day 6")

	// read input
	readFile, err := os.Open("days/day_6/input.txt")
	// readFile, err := os.Open("days/day_6/sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	theMap := [][]string{}

	xStart := -1
	yStart := -1
	dirStart := ""

	directions := []string{UP, DOWN, LEFT, RIGHT}

	lineNo := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		splitLine := strings.Split(line, "")

		theMap = append(theMap, splitLine)

		// only bother searching for the starting position if we haven't yet
		if xStart == -1 {
			i := slices.IndexFunc(splitLine, func(s string) bool { return slices.Contains(directions, s) })
			if i != -1 {
				yStart = lineNo
				xStart = i
				dirStart = splitLine[i]
			}
		}
		lineNo++
	}

	yMax := len(theMap)
	xMax := len(theMap[0])

	readFile.Close()

	fmt.Println("yMax:", yMax)
	fmt.Println("xMax:", xMax)
	fmt.Println("yStart:", yStart)
	fmt.Println("xStart:", xStart)
	fmt.Println("dirStart:", dirStart)

	x := xStart
	y := yStart
	dir := dirStart

	// handle initial position, mark starting position as an X and start at distinctPositions = 1
	distinctPositions := 1
	theMap[yStart][xStart] = "X"

	done := false
	for !done {

		switch dir {
		case LEFT:
			if x-1 < 0 {
				done = true
				break
			}

			if theMap[y][x-1] == OBSTRUCTION {
				dir = UP
				continue
			} else {
				x--
			}

		case RIGHT:
			if x+1 >= xMax {
				done = true
				break
			}

			if theMap[y][x+1] == OBSTRUCTION {
				dir = DOWN
				continue
			} else {
				x++
			}
		case UP:
			if y-1 < 0 {
				done = true
				break
			}

			if theMap[y-1][x] == OBSTRUCTION {
				dir = RIGHT
				continue
			} else {
				y--
			}
		case DOWN:
			if y+1 >= yMax {
				done = true
				break
			}

			if theMap[y+1][x] == OBSTRUCTION {
				dir = LEFT
				continue
			} else {
				y++
			}
		}

		// if visiting a position not visitied before, mark it as an X and count the distinctPosition
		if !done && theMap[y][x] == "." {
			theMap[y][x] = "X"
			distinctPositions++
		}
	}
	fmt.Println("Part 1:", distinctPositions)
}
