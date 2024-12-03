package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// did this one in ts already. do i want to do it here?
func Run() {
	bytes, err := os.ReadFile("./data/d3.txt")
	if err != nil {
		os.Exit(1)
	}
	content := string(bytes)

	content = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	total := 0

	parts := strings.Split(content, "\n")
	for _, part := range parts {
		strParts := strings.Split(part, " ")
		intParts := convertStrToIntSlice(strParts)

		if testReport(intParts) {
			fmt.Println("good", part)
			total += 1
		}
	}

	fmt.Println(total)
}

func convertStrToIntSlice(strParts []string) []int {
	var intParts = make([]int, len(strParts))
	for i, strPart := range strParts {
		x, err := strconv.Atoi(strPart)
		if err != nil {
			os.Exit(1)
		}
		intParts[i] = x
	}
	return intParts
}

// The levels are either all increasing or all decreasingI.
// Any two adjacent levels differ by at least one and at most three.
func testReport(report []int) bool {
	var dir string
	if report[0] > report[1] {
		dir = "desc"
	} else if report[0] < report[1] {
		dir = "inc"
	} else {
		// quick quit
		return false
	}

	println(dir)
	return true
}
