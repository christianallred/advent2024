package day4

import (
	"fmt"
	"os"
)

func Run() {
	bytes, err := os.ReadFile("./data/d4.txt")
	if err != nil {
		os.Exit(1)
	}
	content := string(bytes)
	fmt.Println(content)
}
