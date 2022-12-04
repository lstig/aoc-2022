package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Set map[rune]struct{}
var exists struct{}

func toSet(arr []rune) Set {
	set := Set{}
	for _, r := range arr {
		set[r] = exists
	}
	return set
}

func first(s Set) (*rune, error) {
	for k := range s {
		return &k, nil
	}
	return nil, errors.New("empty set")
}

func priority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r - 38)
	}
	return int(r - 96)
}

// Should make this recursive, e.g. intersect(s1 Set, s2... Set)
func intersect(s1 Set, s2 Set) Set {
	n := Set{}

	for r := range s2 {
		if _, ok := s1[r]; ok {
			n[r] = exists
		}
	}

	return n
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total1 := 0
	total2 := 0
	lines := []Set{}
	n := Set{}

	for scanner.Scan() {

		line := []rune(scanner.Text())

		// Part 1
		n = intersect(toSet(line[:(len(line) / 2)]), toSet(line[len(line) / 2:]))
		for r := range n {
			total1 = total1 + priority(r)
		}

		// Part 2
		lines = append(lines, toSet(line))
		if (len(lines) % 3) == 0 {
			n = intersect(lines[0], intersect(lines[1], lines[2]))
			lines = []Set{}
			for r := range n {
				total2 = total2 + priority(r)
			}
		}

	}

	fmt.Printf("Part 1 %v\n", total1)
	fmt.Printf("Part 2 %v\n", total2)
}