package day5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	bytes, err := os.ReadFile("./data/d5.txt")
	if err != nil {
		os.Exit(1)
	}
	content := string(bytes)

	rules := make([]string, 0)
	reports := make([]string, 0)

	foundBreak := false
	for _, row := range strings.Split(content, "\n") {
		if len(row) == 0 {
			foundBreak = true
			continue
		}
		if foundBreak {
			reports = append(reports, row)
		} else {
			rules = append(rules, row)
		}
	}

	// shoudl we make a list of rules
	beforeRules, afterRules := getBeforeAfterFromRules(rules)

	for key, value := range beforeRules {
		fmt.Println(key, value)
	}

	fmt.Println("-------")

	for key, value := range afterRules {
		fmt.Println(key, value)
	}

	os.Exit(1)
	total := 0
	total2 := 0
	for _, report := range reports {
		parts := strings.Split(report, ",")
		// convert string array to int array
		intReport := make([]int, 0)
		for _, p := range parts {
			part, err := strconv.Atoi(p)
			if err != nil {
				os.Exit(2)
			}
			intReport = append(intReport, part)
		}

		if shouldReportCount(intReport, beforeRules, afterRules) {
			middle := (len(intReport) - 1) / 2
			value := intReport[middle]
			total += value
		} else {
			newReport := reorderReport(intReport, beforeRules, afterRules)

			middle := (len(newReport) - 1) / 2
			value := newReport[middle]
			total2 += value
		}

	}
	fmt.Println("total1: ", total)
	fmt.Println("total2: ", total2)
}

func reorderReport(intReport []int, beforeRules map[int][]int, afterRules map[int][]int) []int {
	// find the ideal range

	// take my numbers, and drop them into it..

	return intReport
}

func shouldReportCount(intReport []int, beforeRules map[int][]int, afterRules map[int][]int) bool {
	for i, part := range intReport {
		beforeParts := intReport[:i]
		afterParts := intReport[i+1:]

		beforeRule, exists := beforeRules[part]
		if exists {
			for _, afterPart := range afterParts {
				if slices.Contains(beforeRule, afterPart) {
					return false
				}
			}
		}

		afterRule, exists := afterRules[part]
		if exists {
			for _, beforePart := range beforeParts {
				if slices.Contains(afterRule, beforePart) {
					return false
				}

			}
		}
	}
	return true
}

func getBeforeAfterFromRules(rules []string) (map[int][]int, map[int][]int) {
	befores := make(map[int][]int)
	afters := make(map[int][]int)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")

		b := parts[0]
		a := parts[1]

		before, err := strconv.Atoi(b)
		if err != nil {
			os.Exit(2)
		}

		after, err := strconv.Atoi(a)
		if err != nil {
			os.Exit(2)
		}

		bValue, exists := befores[after]
		if exists {
			befores[after] = append(bValue, before)
		} else {
			befores[after] = []int{before}
		}

		value, exists := afters[before]
		if exists {
			afters[before] = append(value, after)
		} else {
			afters[before] = []int{after}
		}
	}
	return befores, afters
}

func isManualValid(manual string, rules []string) bool {
	return true
}
