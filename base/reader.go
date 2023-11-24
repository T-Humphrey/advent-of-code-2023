package base

import (
	"bufio"
	"os"
)

var scanner *bufio.Scanner

func LoadFile(path string) {
	if scanner != nil {
		panic("File already loaded, call LoadFile() only once")
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	scanner.Scan()
}

func ReadLine() (string, error) {
	if scanner == nil {
		panic("No file loaded, call LoadFile() first")
	}

	if scanner.Scan() {
		return scanner.Text(), nil
	} else {
		return "", scanner.Err()
	}

}
