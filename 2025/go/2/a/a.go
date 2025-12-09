package a

import (
	"aoc2025/2/util"
	"aoc2025/log"
	"math"
	"strconv"
	"strings"
)

func Solve(input string) int {
	logger := log.NewLogger("[Day 2:a]")
	logger.Println()

	ranges := strings.Split(
		strings.TrimRight(input, "\n"),
		",",
	)

	sum := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		low, err := strconv.Atoi(parts[0])

		if err != nil {
			logger.Fatal(err)
		}

		high, err := strconv.Atoi(parts[1])

		if err != nil {
			logger.Fatal(err)
		}

		logger.Printf("%d - %d:", low, high)

		for id := low; id <= high; id++ {
			idLen := util.IntLen(id)

			if idLen%2 != 0 {
				// only check numbers with even length
				continue
			}

			sig := int(math.Pow(10, float64(idLen/2)))

			pref := id / sig
			aff := id % sig

			if pref == aff {
				sum += id
				logger.Printf("\t%d", id)
			}
		}
	}

	return sum
}
