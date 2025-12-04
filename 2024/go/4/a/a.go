package a

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const (
	North     = 010000000
	NorthEast = 001000000
	East      = 000001000
	SouthEast = 000000001
	South     = 000000010
	SouthWest = 000000100
	West      = 000100000
	NorthWest = 100000000
)

func SolvePartA() int {
	file, err := os.Open("./input")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	occurrences := 0

	for lineIndex, line := range lines {
		regex := regexp.MustCompile("X")
		results := regex.FindAllStringIndex(line, -1)

		for _, result := range results {
			charIndex := result[0]

			findM(lines, lineIndex, charIndex)
		}

		fmt.Println()
	}

	return occurrences
}

func findM(lines []string, lineIndex int, charIndex int) []int {
	matches := make([]int, 0)

	minLine := lineIndex
	if lineIndex >= 3 {
		minLine = lineIndex - 1
	}

	maxLine := lineIndex
	if lineIndex <= len(lines)-4 {
		maxLine = lineIndex + 1
	}

	minChar := charIndex
	if charIndex >= 3 {
		minChar = charIndex - 1
	}

	maxChar := charIndex
	if charIndex <= len(lines[lineIndex])-4 {
		maxChar = charIndex + 1
	}

	for l := minLine; l <= maxLine; l++ {
		for c := minChar; c <= maxChar; c++ {
			char := lines[l][c : c+1]
			if "M" == char {

				fmt.Print()
			}
			fmt.Print(char)
		}
		fmt.Printf("\n")
	}

	return matches
}
