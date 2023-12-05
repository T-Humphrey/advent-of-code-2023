package day_4

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/t-humphrey/advent-of-code-2023/base"
)

type card struct {
	number  int
	nums    map[int]int
	winners map[int]int
}

var mem = make(map[int]int)

func RunPart1() {
	lines := base.LoadFile("day_4/input.txt")
	runningTotal := 0
	for _, line := range lines {
		game := parse(line)
		winners := numWinners(game)
		runningTotal += int(math.Pow(2, float64(winners)-1))
	}

	fmt.Println(runningTotal)
}

func RunPart2() {
	lines := base.LoadFile("day_4/input.txt")
	var games []card
	for _, line := range lines {
		games = append(games, parse(line))
	}

	count := 0
	for i := 0; i < len(games); i++ {
		count += run(games, i)
	}
	fmt.Println(count)
}

func run(games []card, i int) int {
	if val, ok := mem[i]; ok {
		return val
	}
	num := numWinners(games[i])
	if num == 0 {
		return 1
	}

	count := 0
	for j := 1; j <= num; j++ {
		count += run(games, i+j)
	}
	mem[i] = count + 1
	return count + 1
}

func numWinners(game card) int {
	count := 0
	for n, _ := range game.nums {
		if _, ok := game.winners[n]; ok {
			count++
		}
	}
	return count
}

func parse(line string) card {
	nums := make(map[int]int)
	winners := make(map[int]int)

	number := strings.TrimSpace(line[4:8])
	iNumber, _ := strconv.Atoi(number)
	l := strings.Split(line, ":")[1]
	sides := strings.Split(l, "|")
	for _, v := range strings.Split(sides[0], " ") {
		if strings.TrimSpace(v) == "" {
			continue
		}

		n, _ := strconv.Atoi(v)
		winners[n]++
	}
	for _, v := range strings.Split(sides[1], " ") {
		if strings.TrimSpace(v) == "" {
			continue
		}

		n, _ := strconv.Atoi(v)
		nums[n]++
	}
	return card{
		number:  iNumber,
		nums:    nums,
		winners: winners,
	}
}
