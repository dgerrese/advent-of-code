package b

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolvePartB() int {
	file, err := os.Open("./input")
	left := make([]int, 0)
	right := make([]int, 0)

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

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		leftNumber, convErr := strconv.Atoi(numbers[0])
		if convErr != nil {
			log.Fatal(err)
		}
		rightNumber, convErr := strconv.Atoi(numbers[1])
		if convErr != nil {
			log.Fatal(err)
		}

		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	similarity := 0
	for i := 0; i < len(left); i++ {
		number := left[i]
		occurences := 0
		for j := 0; j < len(right); j++ {
			if number == right[j] {
				occurences++
			}
		}
		similarity += number * occurences
	}

	return similarity
}
