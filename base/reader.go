package base

import (
	"bufio"
	"os"
)

var scanner *bufio.Scanner

func LoadFile(path string) []string {
	if scanner != nil {
		panic("File already loaded, call LoadFile() only once")
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)
	scanner.Scan()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}
