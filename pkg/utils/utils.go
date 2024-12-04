package utils

import (
	"strings"
)

// this func auto drops empty rows
func MakeTableWithData(data string, rowSep string, colSep string) [][]string {
	var table [][]string
	for _, row := range strings.Split(data, rowSep) {
		if len(row) > 0 {
			table = append(table, strings.Split(row, colSep))
		}
	}
	return table
}
