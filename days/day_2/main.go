package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")

	// read input
	// readFile, err := os.Open("days/day_2/input.txt")
	readFile, err := os.Open("days/day_2/2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	reports := [][]int{}

	for fileScanner.Scan() {
		report := []int{}
		levelsStrings := strings.Split(fileScanner.Text(), " ")
		for _, levelString := range levelsStrings {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				panic("yo, cant even read the input go home")
			}
			report = append(report, level)
		}
		reports = append(reports, report)
	}
	readFile.Close()

	part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		increasing := true
		decreasing := true
		differRange := true
		for i := 1; i < len(report); i++ {
			decreasing = decreasing && report[i] < report[i-1]
			increasing = increasing && report[i] > report[i-1]
			differRange = differRange && differsBy1To3(report[i], report[i-1])
		}
		if (increasing || decreasing) && differRange {
			safeReports++
		}
	}

	fmt.Println("Part 1:", safeReports)
}

func differsBy1To3(num1, num2 int) bool {
	num := num1 - num2
	if num < 0 {
		num *= -1
	}
	return num >= 1 && num <= 3
}

func part2(reports [][]int) {
	safeReports := 0
	for i, report := range reports {
		evaluation := evaluateReport(report)

		fmt.Println("\nLine", i)
		fmt.Println("report:", report)
		fmt.Println("evaluation:", evaluation)

		if evaluation.isSafe {
			fmt.Println("safe by default:", report)
			safeReports++
			continue
		}

		// report isn't safe - handle cases where we see we can become safe by removing one level

		// First handle cases where one increasing/decreasing differs from the rest can be fixed.
		// For these cases, remove the offender and run it through again to see if it satisfies the
		// differs by rule, since it should pass the increasing/decreasing now.
		if len(evaluation.decreases) == 1 && len(evaluation.increases) == len(report)-2 && len(evaluation.differViolations) < 2 {
			// fmt.Println(i)
			fmt.Println("!!! found one with only 1 decrease:", report)
			shortReport := append(report[:evaluation.decreases[0]], report[evaluation.decreases[0]+1:]...)
			shortReportEvaluation := evaluateReport(shortReport)
			fmt.Println("ew report:", shortReport)
			if shortReportEvaluation.isSafe {
				fmt.Println("got fixed!")
				safeReports++
				continue
			}
		}

		if len(evaluation.increases) == 1 && len(evaluation.decreases) == len(report)-2 && len(evaluation.differViolations) < 2 {
			fmt.Println("@@@ found one with only 1 increase:", report)
			shortReport := append(report[:evaluation.increases[0]], report[evaluation.increases[0]+1:]...)
			shortReportEvaluation := evaluateReport(shortReport)
			fmt.Println("short", shortReport)
			if shortReportEvaluation.isSafe {
				fmt.Println("got fixed")
				safeReports++
				continue
			}
		}

		// handle if there is exactly 1 pair of duplicate numbers, then removing one might be enough to fix
		if len(evaluation.duplicates) == 1 {
			fmt.Println("found dupe:", report)
			shortReport := append(report[:evaluation.duplicates[0]], report[evaluation.duplicates[0]+1:]...)
			shortReportEvaluation := evaluateReport(shortReport)
			fmt.Println("short", shortReport)
			if shortReportEvaluation.isSafe {
				fmt.Println("got fixed")
				safeReports++
				continue
			}
		}

		// next handle cases where there are either 1 or 2 calcualtions that resulted in a violation of the diff rule, since
		// this either 1 or 2 cases could be casued by a single number
		if len(evaluation.differViolations) == 2 {
			fmt.Println("### found one diff fixable maybe:", report)
			shortReport1 := append(report[:evaluation.differViolations[len(evaluation.differViolations)-1]], report[evaluation.differViolations[len(evaluation.differViolations)-1]+1:]...)
			shortReportEvaluation1 := evaluateReport(shortReport1)
			fmt.Println("short", shortReport1)
			if shortReportEvaluation1.isSafe {
				fmt.Println("got fixed")
				safeReports++
				continue
			}

			fmt.Println("### found one diff fixable maybe:", report)
			shortReport2 := append(report[:evaluation.differViolations[len(evaluation.differViolations)-2]], report[evaluation.differViolations[len(evaluation.differViolations)-2]+1:]...)
			shortReportEvaluation2 := evaluateReport(shortReport2)
			fmt.Println("short", shortReport2)
			if shortReportEvaluation2.isSafe {
				fmt.Println("got fixed")
				safeReports++
				continue
			}

		}

		if len(evaluation.differViolations) == 1 {
			fmt.Println("### found one diff fixable maybe:", report)
			shortReport := append(report[:evaluation.differViolations[len(evaluation.differViolations)-1]], report[evaluation.differViolations[len(evaluation.differViolations)-1]+1:]...)
			shortReportEvaluation := evaluateReport(shortReport)
			fmt.Println("short", shortReport)
			if shortReportEvaluation.isSafe {
				fmt.Println("got fixed")
				safeReports++
				continue
			}
		}

		fmt.Println("  not fixable:", report)
		fmt.Println("  violations:", evaluation.violationCount)
		fmt.Println("  eval:", evaluation)
	}
	fmt.Println("Part 2:", safeReports)
}

type ReportEvaluation struct {
	isSafe           bool
	increasing       bool
	decreasing       bool
	increases        []int // this indexes previous element was less than it
	decreases        []int // this indexes previous element was greater than it
	differViolations []int
	duplicates       []int

	violationCount int
}

// evaluateReport returns isSafe
func evaluateReport(report []int) ReportEvaluation {
	allIncreasing := true
	allDecreasing := true
	decreases := []int{}
	increases := []int{}
	differRange := true
	differViolations := []int{}
	duplicates := []int{}

	for i := 1; i < len(report); i++ {
		decreasing := report[i] < report[i-1]
		if decreasing {
			decreases = append(decreases, i)
		}
		allDecreasing = allDecreasing && decreasing

		increasing := report[i] > report[i-1]
		if increasing {
			increases = append(increases, i)
		}
		allIncreasing = allIncreasing && increasing

		if report[i] == report[i-1] {
			duplicates = append(duplicates, i)
		}

		differsOK := differsBy1To3(report[i], report[i-1])

		if !differsOK {
			differViolations = append(differViolations, i)
		}

		differRange = differRange && differsOK
	}

	isSafe := (allIncreasing || allDecreasing) && differRange

	violationCount := 0
	if !allIncreasing && !allDecreasing {
		violationCount++
	}
	if len(duplicates) > 0 {
		violationCount += len(duplicates)
	}
	if len(differViolations) > 0 {
		violationCount += len(differViolations)
	}

	return ReportEvaluation{
		isSafe:           isSafe,
		increasing:       allIncreasing,
		decreasing:       allDecreasing,
		increases:        increases,
		decreases:        decreases,
		differViolations: differViolations,
		duplicates:       duplicates,
		violationCount:   violationCount,
	}
}

// bug

// run as part of big run
// ...
// not fixable: [30 30 28 27 26 23 26 26]
// violations: 5
// eval: {false false false [7] [2 3 4 6] [1 5] [1 5] 5}

// not fixable: [40 36 30 29 29]
// violations: 2
// eval: {false false true [] [1 2 3 4] [1 2] [] 2}

// not fixable: [59 59 62 65 65 66 66]
// violations: 7
// eval: {false false false [2 3 5] [] [1 4 6] [1 4 6] 7}
// ...

// run alone
// not fixable: [40 36 30 29 29]
// violations: 5
// eval: {false false false [] [1 2 3] [1 2 4] [4] 5}
