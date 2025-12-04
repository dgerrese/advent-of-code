package a

import (
	"strconv"
	"strings"
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	dial := 50
	hits := 0

	for _, l := range lines {
		if len(l) < 2 {
			continue
		}

		dir := l[:1]
		dist, err := strconv.Atoi(l[1:])

		if err != nil {
			panic("Could not parse line: " + l)
		}

		switch dir {
		case "R":
			dial = (dial + dist) % 100

			if dial == 0 {
				hits++
			}
		case "L":
			dial = (dial - dist) % 100

			if dial == 0 {
				hits++
			}
		}
	}

	return hits
}
