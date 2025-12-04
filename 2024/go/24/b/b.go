package b

import (
	"bufio"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

type Statement struct {
	left  string
	gate  string
	right string
	out   string
}

func SolvePartB() int {
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

	statements := make([]Statement, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "->") {
			parts := strings.Split(line, " ")
			statements = append(statements, Statement{parts[0], parts[1], parts[2], parts[4]})
		}
	}

	slices.SortFunc(statements, func(a Statement, b Statement) int {
		return strings.Compare(a.out, b.out)
	})

	wires := make(map[string][]string)

	for _, statement := range statements {
		log.Println(statement)

		key := statement.out[0:1]

		wires[key] = append(wires[key], statement.out+"("+statement.left+","+statement.right+")")
	}

	for wire, parts := range wires {
		sort.Strings(parts)
		parts = slices.Compact(parts)
		log.Println(wire, parts)
	}

	return 0
}

//func findStatementValue(inputs map[string]int, statements []Statement, statement Statement) int {
//	left := -1
//	right := -1
//
//	for input, value := range inputs {
//		if input == statement.left {
//			left = value
//			continue
//		}
//		if input == statement.right {
//			right = value
//			continue
//		}
//	}
//
//	if left == -1 {
//		for _, statement2 := range statements {
//			if statement.left == statement2.out {
//				left = findStatementValue(inputs, statements, statement2)
//				break
//			}
//		}
//	}
//
//	if right == -1 {
//		for _, statement2 := range statements {
//			if statement.right == statement2.out {
//				right = findStatementValue(inputs, statements, statement2)
//				break
//			}
//		}
//	}
//
//	if statement.gate == "OR" {
//		return left | right
//	}
//
//	if statement.gate == "AND" {
//		return left & right
//	}
//
//	return left ^ right
//}
