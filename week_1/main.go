package week_1

import "github.com/t-humphrey/advent-of-code-2023/base"

func Run() {
	base.LoadFile("week_1/input.txt")
	for {
		line, err := base.ReadLine()
		if err != nil {
			break
		}
		println(line)
	}
}
