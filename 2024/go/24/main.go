package main

import (
	"day24/a"
	"day24/b"
	"fmt"
)

func main() {
	sa := a.SolvePartA()
	sb := b.SolvePartB()

	fmt.Printf("A: %d\n", sa)
	fmt.Printf("B: %d\n", sb)
}
