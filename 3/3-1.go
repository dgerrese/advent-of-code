package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
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

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`mul\((?P<arg1>\d+),(?P<arg2>\d+)\)`)
		results := regex.FindAllStringSubmatch(line, -1)

		for _, result := range results {
			number1, _ := strconv.Atoi(result[1])
			number2, _ := strconv.Atoi(result[2])

			sum += number1 * number2
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
