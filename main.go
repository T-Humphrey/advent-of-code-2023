package main

import (
	"os"

	"github.com/t-humphrey/advent-of-code-2023/week_1"
)

func main() {
	week := os.Args[1]
	switch week {
	case "week_1":
		week_1.Run()
	}
}
