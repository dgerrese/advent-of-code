package day2

import (
	"aoc2025/2/a"
	"aoc2025/2/b"
	_ "embed"
	"fmt"
)

var (
	//go:embed input
	input string
)

func Solve() {
	sa := a.Solve(input)
	sb := b.Solve(input)

	fmt.Printf("A:\t%d\n", sa)
	fmt.Printf("B:\t%d\n", sb)
}
