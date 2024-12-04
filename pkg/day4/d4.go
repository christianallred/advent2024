package day4

import (
	"fmt"
	"os"
	"playground/pkg/utils"
	"slices"
	"strings"
)

func Run() {
	bytes, err := os.ReadFile("./data/d4.txt")
	if err != nil {
		os.Exit(1)
	}
	content := string(bytes)
	// 	content = `MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX`

	// content = `AM.S.....A
	// ..A..MSMS.
	// .M.S.MAA..
	// ..A.ASMSM.
	// .M.S.M....
	// ..........
	// S.S.S.S.S.
	// .A.A.A.A..
	// M.M.M.M.M.
	// A........A`

	table := utils.MakeTableWithData(content, "\n", "")
	xmas1 := 0
	xmas2 := 0
	for y, row := range table {
		for x, col := range row {
			if col == "X" {
				xmas1 += checkXmas1(y, x, table)
			}
			if col == "A" {
				// check to makes ure its at least 1 from every edge
				if x >= 1 && len(table[0])-x > 1 && y >= 1 && len(table)-y > 1 {
					xmas2 += checkXmas2(y, x, table)
				}
			}
		}
	}

	fmt.Println("xmas1: ", xmas1)
	fmt.Println("xmas2: ", xmas2)
}

func checkXmas2(y int, x int, table [][]string) int {
	// M.S
	// .A.
	// M.S
	//
	// S.M
	// .A.
	// S.M
	//
	// M.M
	// .A.
	// S.S
	//
	// S.S
	// .A.
	// M.M

	patterns := []string{"MSSM", "SMMS", "MMSS", "SSMM"}
	positions := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	// get the pattern from this A
	s := make([]string, 4)
	for i, pos := range positions {
		realY := y + pos[0]
		realX := x + pos[1]
		s[i] = table[realY][realX]
	}
	joined := strings.Join(s, "")

	if slices.Contains(patterns, joined) {
		return 1
	}

	return 0
}
func checkXmas1(y int, x int, table [][]string) int {
	maxY := len(table)
	maxX := len(table[0])

	xmas := 0

	vN := [3][2]int{{-1, 0}, {-2, 0}, {-3, 0}}
	vS := [3][2]int{{1, 0}, {2, 0}, {3, 0}}
	hW := [3][2]int{{0, -1}, {0, -2}, {0, -3}}
	hE := [3][2]int{{0, 1}, {0, 2}, {0, 3}}
	dNE := [3][2]int{{-1, 1}, {-2, 2}, {-3, 3}}
	dSE := [3][2]int{{1, 1}, {2, 2}, {3, 3}}
	dSW := [3][2]int{{1, -1}, {2, -2}, {3, -3}}
	dNW := [3][2]int{{-1, -1}, {-2, -2}, {-3, -3}}

	// has padding right
	padRight := maxX-x >= 4
	padLeft := x > 3
	padBelow := maxY-y >= 4
	padAbove := y > 3

	// Horizontal
	if padRight {
		xmas += checkXmas(table, y, x, hE)
	}
	if padLeft {
		xmas += checkXmas(table, y, x, hW)
	}
	// Vertical
	if padBelow {
		xmas += checkXmas(table, y, x, vS)
	}
	if padAbove {
		xmas += checkXmas(table, y, x, vN)
	}
	// Diaganol
	if padLeft && padAbove {
		xmas += checkXmas(table, y, x, dNW)
	}
	if padRight && padAbove {
		xmas += checkXmas(table, y, x, dNE)
	}
	if padRight && padBelow {
		xmas += checkXmas(table, y, x, dSE)
	}
	if padLeft && padBelow {
		xmas += checkXmas(table, y, x, dSW)
	}
	return xmas
}

func checkXmas(table [][]string, y int, x int, dir [3][2]int) int {

	for i, row := range dir {
		realY := y + row[0]
		realX := x + row[1]
		value := table[realY][realX]

		if i == 0 && value != "M" {
			return 0
		}
		if i == 1 && value != "A" {
			return 0
		}
		if i == 2 && value != "S" {
			return 0
		}
	}
	return 1
}
