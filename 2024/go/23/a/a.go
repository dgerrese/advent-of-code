package a

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
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

	computers := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		lineComputers := strings.Split(line, "-")

		leftComputer := lineComputers[0]
		rightComputer := lineComputers[1]

		computers[leftComputer] = append(computers[leftComputer], rightComputer)
		computers[rightComputer] = append(computers[rightComputer], leftComputer)
	}

	groups := make([]string, 0)

	for computer, connectedComputers := range computers {
		if !("t" == computer[0:1]) {
			continue
		}

		sort.Strings(connectedComputers)
		for _, connectedComputer := range connectedComputers {
			subComputers := computers[connectedComputer]

			sort.Strings(subComputers)
			for _, subComputer := range subComputers {
				subSubComputers := computers[subComputer]

				sort.Strings(subSubComputers)
				for _, subSubComputer := range subSubComputers {
					if subSubComputer == computer {
						groupComputers := []string{computer, connectedComputer, subComputer}
						sort.Strings(groupComputers)
						group := strings.Join(groupComputers, ",")

						dupe := false
						for _, g := range groups {
							if g == group {
								dupe = true
							}
						}

						if !dupe {
							groups = append(groups, group)
						}
					}
				}
			}
		}
	}

	return len(groups)
}
