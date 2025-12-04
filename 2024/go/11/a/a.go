package a

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(fileName string, delim byte, handler func(line string)) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString(delim)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file line error: %v", err)
			return
		}
		handler(line)
	}
}

func SolvePartA() int {
	numbers := []string{}

	readInput("./input", '\n', func(line string) {
		numbers = strings.Split(strings.TrimSuffix(line, "\n"), " ")
	})

	for i := 0; i < 75; i++ {
		//out := []string{}
		for _, number := range numbers {
			mapNumber(number)
		}
	}

	return 0
}

func mapNumber(num string) []string {
	if num == "0" {
		return []string{"1"}
	}

	if len(num)%2 == 0 {
		out := []string{}

		out = append(out, num[0:len(num)/2])
		out = append(out, num[len(num)/2:])

		return out
	}

	atoi, _ := strconv.Atoi(num)
	return []string{strconv.Itoa(atoi * 2024)}
}
