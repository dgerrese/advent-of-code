package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Statement struct {
	left  string
	gate  string
	right string
	out   string
}

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

	inputs := make(map[string]int)
	statements := make([]Statement, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "->") {
			parts := strings.Split(line, " ")
			statements = append(statements, Statement{parts[0], parts[1], parts[2], parts[4]})
		} else if "" == line {
			continue
		} else {
			parts := strings.Split(line, " ")
			value, _ := strconv.Atoi(parts[1])
			inputs[parts[0][0:3]] = value
		}
	}

	outputs := make(map[string]int)

	for _, statement := range statements {
		if statement.out[0:1] == "z" {
			value := findStatementValue(inputs, statements, statement)
			outputs[statement.out] = value
		}
	}

	outputKeys := make([]string, 0)

	for key := range outputs {
		outputKeys = append(outputKeys, key)
	}
	sort.Strings(outputKeys)
	slices.Reverse(outputKeys)

	binaryOutput := ""

	for _, key := range outputKeys {
		binaryOutput += strconv.Itoa(outputs[key])
	}

	output, _ := strconv.ParseInt(binaryOutput, 2, 64)
	fmt.Printf("Output for wires starting with z: %v\n", output)
}

func findStatementValue(inputs map[string]int, statements []Statement, statement Statement) int {
	left := -1
	right := -1

	for input, value := range inputs {
		if input == statement.left {
			left = value
			continue
		}
		if input == statement.right {
			right = value
			continue
		}
	}

	if left == -1 {
		for _, statement2 := range statements {
			if statement.left == statement2.out {
				left = findStatementValue(inputs, statements, statement2)
				break
			}
		}
	}

	if right == -1 {
		for _, statement2 := range statements {
			if statement.right == statement2.out {
				right = findStatementValue(inputs, statements, statement2)
				break
			}
		}
	}

	if statement.gate == "OR" {
		return left | right
	}

	if statement.gate == "AND" {
		return left & right
	}

	return left ^ right
}
