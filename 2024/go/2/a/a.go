package a

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

	safeReports := 0
	for scanner.Scan() {
		report := scanner.Text()

		levels := strings.Split(report, " ")

		trend := 0
		isSafe := true

		for i := 1; i < len(levels); i++ {
			level, err := strconv.Atoi(levels[i])
			if err != nil {
				log.Fatal(err)
			}
			previous, err := strconv.Atoi(levels[i-1])
			if err != nil {
				log.Fatal(err)
			}

			if previous == level {
				isSafe = false
				break
			}

			if previous < level {
				if trend < 0 || (level-previous > 3) {
					isSafe = false
					break
				}
				trend = 1
			}

			if previous > level {
				if trend > 0 || (previous-level > 3) {
					isSafe = false
					break
				}
				trend = -1
			}
		}

		if isSafe {
			safeReports++
		}
	}

	return safeReports
}
