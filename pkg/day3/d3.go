package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// TODO: a fun challenge woudl be to iterate through this in real time as we read teh bytes out
func Run() {
	bytes, err := os.ReadFile("./data/d3.txt")
	if err != nil {
		os.Exit(1)
	}
	content := string(bytes)

	doParts := strings.Split(content, "do()")

	total := 0
	for _, doPart := range doParts {
		actualDo := strings.Split(doPart, "don't()")
		if len(actualDo) > 0 {
			total += calcPartTotal(actualDo[0])
		}
	}
	fmt.Println("part1: ", calcPartTotal(content))
	fmt.Println("part2: ", total)
}

func calcPartTotal(t string) int {
	total := 0
	parts := strings.Split(t, "mul(")
	for _, part := range parts {
		bPart := []byte(part)
		exp, err := regexp.Compile(`^\d{1,3},\d{1,3}\)`)

		if err != nil {
			fmt.Println("invalid regex")
			os.Exit(1)
		}

		found := exp.Find(bPart)
		if len(found) > 0 {
			foundStripped := strings.Replace(string(found), ")", "", 1)
			numParts := strings.Split(foundStripped, ",")

			x, err := strconv.Atoi(numParts[0])
			if err != nil {
				fmt.Println("error parsting x")
				os.Exit(1)
			}
			y, err := strconv.Atoi(numParts[1])
			if err != nil {
				fmt.Println("error parsting y")
				os.Exit(1)
			}
			total += x * y
		}
	}
	return total
}
