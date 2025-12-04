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

		if validateReport(levels, true) {
			safeReports++
		}
	}

	return safeReports
}

func removeIndex(s []string, index int) []string {
	result := make([]string, 0)
	result = append(result, s[:index]...)
	return append(result, s[index+1:]...)
}

func validateReport(levels []string, withDampening bool) bool {
	trend := 0

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
			if withDampening {
				return validateReport(removeIndex(levels, i), false)
			}

			return false
		}

		if previous < level {
			if trend < 0 || (level-previous > 3) {
				if withDampening {
					return validateReport(removeIndex(levels, i), false)
				}

				return false
			}
			trend = 1
		}

		if previous > level {
			if trend > 0 || (previous-level > 3) {
				if withDampening {
					return validateReport(removeIndex(levels, i), false)
				}

				return false
			}
			trend = -1
		}
	}

	return true
}
