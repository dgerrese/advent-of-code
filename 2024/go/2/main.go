package main

import (
	"day2/a"
	"day2/b"
	"fmt"
)

func main() {
	sa := a.SolvePartA()
	sb := b.SolvePartB()

	fmt.Printf("A: %d\n", sa)
	fmt.Println()
	fmt.Printf("B: %d\n", sb)
}
