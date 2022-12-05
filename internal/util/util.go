package util

import (
	"bufio"
	"log"
	"os"
)

func OpenInput(input string) *bufio.Scanner {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln(err)
	}

	return bufio.NewScanner(file)
}