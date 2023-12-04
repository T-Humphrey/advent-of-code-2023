package day_2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/t-humphrey/advent-of-code-2023/base"
)

type Game struct {
	Number int
	Rounds []Round
}

type Round struct {
	Cubes []Cube
}

type Cube struct {
	Colour string
	Count  int
}

var cubeCount = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func RunPart1() {
	lines := base.LoadFile("day_2/input.txt")
	runningTotal := 0

	for _, line := range lines {
		game := parseLine(line)
		invalid := false
		for _, round := range game.Rounds {
			if !validRound(round) {
				invalid = true
			}
		}
		if !invalid {
			runningTotal += game.Number
		}

	}
	fmt.Println(runningTotal)
}

func RunPart2() {
	lines := base.LoadFile("day_2/input.txt")
	runningTotal := 0

	for _, line := range lines {
		game := parseLine(line)
		minNeeded := make(map[string]int)
		for _, r := range game.Rounds {
			for _, c := range r.Cubes {
				v, ok := minNeeded[c.Colour]
				if ok {
					minNeeded[c.Colour] = int(math.Max(float64(v), float64(c.Count)))
				} else {
					minNeeded[c.Colour] = c.Count
				}
			}
		}
		pow := minNeeded["blue"] * minNeeded["red"] * minNeeded["green"]
		runningTotal += pow
	}
	fmt.Println(runningTotal)
}

func validRound(round Round) bool {
	for _, cube := range round.Cubes {
		if cube.Count > cubeCount[cube.Colour] {
			return false
		}
	}
	return true
}

func parseLine(line string) Game {
	gameSplit := strings.Split(line, ":")
	gameNumber, _ := strconv.Atoi(strings.Split(gameSplit[0], " ")[1])
	roundStrings := strings.Split(gameSplit[1], ";")
	var rounds []Round
	for _, r := range roundStrings {
		parts := strings.Split(r, ",")
		var cubes []Cube
		for _, p := range parts {
			s := strings.Split(strings.TrimSpace(p), " ")
			cnt, _ := strconv.Atoi(s[0])
			c := Cube{
				Colour: s[1],
				Count:  cnt,
			}
			cubes = append(cubes, c)
		}
		round := Round{
			Cubes: cubes,
		}
		rounds = append(rounds, round)
	}
	return Game{
		Number: gameNumber,
		Rounds: rounds,
	}
}
