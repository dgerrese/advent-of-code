package day1

import (
	"aoc2025/1/a"
	"aoc2025/1/b"
	_ "embed"
	"fmt"
	"log"
)

var (
	//go:embed input
	input string
)

func Solve(log *log.Logger) {
	log.Println("Log test")
	sa := a.Solve(input)
	ba := b.Solve(input)

	fmt.Printf("A:\t%d\n", sa)
	fmt.Printf("B:\t%d\n", ba)
}
