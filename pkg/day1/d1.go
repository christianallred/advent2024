package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("day1 start")
	file, err := os.Open("./data/d1.txt")
	if err != nil {
		fmt.Println("could not open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a []int
	var b []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		apart, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error converting", parts[0])
		}

		bpart, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("error converting", parts[1])
		}

		a = append(a, apart)
		b = append(b, bpart)
		fmt.Println(line, parts)
	}

	sort.Ints(a)
	sort.Ints(b)

	total := 0
	for i, apart := range a {
		bpart := b[i]
		total += intAbs(apart - bpart)
	}

	fmt.Println(total)
}

func intAbs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
