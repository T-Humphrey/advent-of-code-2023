package day_3

import (
	"fmt"
	"strconv"

	"github.com/t-humphrey/advent-of-code-2023/base"
)

type Part struct {
	Number      int
	Length      int
	Coordinates coordinates
}

type coordinates struct {
	X int
	Y int
}

func RunPart1() {
	lines := base.LoadFile("day_3/input.txt")
	runningTotal := 0
	parts, symbols := parseDoc(lines)
	for _, p := range parts {
		if findSym(p, symbols) {
			runningTotal += p.Number
		}
	}

	fmt.Println(runningTotal)
}

func findSym(part Part, coordinates []coordinates) bool {
	startIdx := part.Coordinates.X - 1
	endIdx := part.Coordinates.X + part.Length + 1
	for _, c := range coordinates {

		if (part.Coordinates.Y == (c.Y-1) || part.Coordinates.Y == c.Y || part.Coordinates.Y == (c.Y+1)) && (c.X >= startIdx && c.X <= endIdx) {
			// fmt.Printf("FOUND- %d was at %d,%d. Symbol at %d,%d\n", part.Number, part.Coordinates.X, part.Coordinates.Y, c.X, c.Y)
			// fmt.Printf("Checked from %d to %d\n", startIdx, endIdx)
			// fmt.Printf("Length %d\n", part.Length)
			return true
		}

	}
	return false
}

func parseDoc(lines []string) ([]Part, []coordinates) {
	var parts []Part
	var symbols []coordinates
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if isNumber(line[j]) {
				num := string(line[j])
				for k := j + 1; k < len(line); k++ {
					if isNumber(line[k]) {
						num += string(line[k])
					} else {
						j = k
						break
					}
				}
				iNum, _ := strconv.Atoi(num)
				p := Part{
					Number: iNum,
					Length: len(num),
					Coordinates: coordinates{
						X: j - len(num),
						Y: i,
					},
				}
				parts = append(parts, p)
			} else if isSymbol(line[j]) {
				symbols = append(symbols, coordinates{
					X: j,
					Y: i,
				})
			}
		}
	}
	return parts, symbols
}

func isNumber(r byte) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r byte) bool {
	return r != '.' && !isNumber(r)
}
