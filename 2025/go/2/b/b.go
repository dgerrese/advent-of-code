package b

import (
	"aoc2025/2/util"
	"aoc2025/log"
	"math"
	"strconv"
	"strings"
)

func Solve(input string) int {
	logger := log.NewLogger("[Day 2:b]")
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

			for pLen := 1; pLen <= idLen/2; pLen++ {
				if idLen%pLen != 0 {
					continue
				}
				logger.Printf("\t%d, pLen: %d", id, pLen)
				segCt := idLen / pLen
				sig := int(math.Pow(10, float64(pLen)))

				pattern := id % sig
				lower := pattern * int(math.Pow(10, float64(segCt-1)))
				higher := (pattern + 1) * int(math.Pow(10, float64(segCt-1)))
				logger.Printf("\t%d [%d] %d, %d\n", id, pattern, lower, higher)
				if pattern == 0 || id < lower || id > higher {
					continue
				}

				match := 0
				for i := 0; i < segCt; i++ {
					match += pattern * int(math.Pow(float64(sig), float64(i)))
				}

				logger.Printf(
					"\tid: %d, idLen: %d, pLen: %d, segCt: %d, sig: %d, pattern: %d, match: %d",
					id,
					idLen,
					pLen,
					segCt,
					sig,
					pattern,
					match,
				)

				if match == id {
					sum += id
					logger.Printf("\t%d", id)
					break
				}
			}
		}
	}

	return sum
}
