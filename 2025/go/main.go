package main

import (
	day1 "aoc2025/1"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Advent of Code 2025")

	// Day 1
	fmt.Printf("\n\nDay 1\n")
	day1.Solve(getLogger("[Day 1] "))

	// Day 2
	// ...
}

func getLogger(pf string) *log.Logger {
	f, err := os.OpenFile("aoc_2025.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)

	if err != nil {
		panic(err)
	}

	return log.New(f, pf, log.Ltime)
}
