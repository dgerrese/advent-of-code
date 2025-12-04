package b

import (
	"math"
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

		ndial := 0

		switch dir {
		case "R":
			ndial = dial + dist
		case "L":
			ndial = dial - dist
		}

		if ndial == 0 {
			hits++
		} else if ndial < 0 {
			revs := math.Abs(float64(ndial) / 100)
			add := 0
			if dial > 0 {
				add++
			}

			ndial += int(math.Ceil(revs)) * 100
			hits += add + int(math.Floor(revs))

		} else if ndial >= 100 {
			revs := math.Abs(float64(ndial) / 100)

			ndial -= int(math.Floor(revs)) * 100
			hits += int(math.Floor(revs))
		}

		dial = ndial
	}

	return hits
}
