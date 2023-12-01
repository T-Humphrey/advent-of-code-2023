package week_1

import (
	"fmt"
	"strconv"

	"github.com/t-humphrey/advent-of-code-2023/base"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func RunPart1() {
	lines := base.LoadFile("week_1/input.txt")
	runingTotal := 0
	for _, line := range lines {
		first := ""
		last := ""

		for i := 0; i < len(line); i++ {
			char := line[i]
			if isNum(char) {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			}
		}
		amount, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		runingTotal += amount
	}
	fmt.Println(runingTotal)
}

func RunPart2() {
	lines := base.LoadFile("week_1/input.txt")
	runingTotal := 0
	for _, line := range lines {
		first := ""
		last := ""

		for i := 0; i < len(line); i++ {
			char := line[i]
			if isNum(char) {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			} else {
				num, ok := tryFindNumber(line[i:])
				if ok {
					if first == "" {
						first = num
					}
					last = num
				}
			}
		}
		amount, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		runingTotal += amount
	}
	fmt.Println(runingTotal)
}

func isNum(char byte) bool {
	return char >= 48 && char <= 57
}

func tryFindNumber(word string) (string, bool) {
	if len(word) >= 3 {
		if num, ok := numMap[word[:3]]; ok {
			return num, true
		}
	}
	if len(word) >= 4 {
		if num, ok := numMap[word[:4]]; ok {
			return num, true
		}
	}
	if len(word) >= 5 {
		if num, ok := numMap[word[:5]]; ok {
			return num, true
		}
	}
	return "", false
}
