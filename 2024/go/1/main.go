package main

import (
	"day1/a"
	"day1/b"
	"fmt"
)

func main() {
	sa := a.SolvePartA()
	sb := b.SolvePartB()

	fmt.Printf("A: %d\n", sa)
	fmt.Println()
	fmt.Printf("B: %d\n", sb)
}
