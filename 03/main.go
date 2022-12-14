package main

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/lstig/aoc-2022/internal/util"
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

func intersect(s1 Set, s2 ...Set) Set {
	n := Set{}
	this, next := s2[0], s2[1:]

	for r := range this {
		if _, ok := s1[r]; ok {
			n[r] = exists
		}
	}

	if len(next) > 0 {
		n = intersect(n, next...)
	}

	return n
}

func main() {
	scanner := util.OpenInput("./input.txt")
	total1 := 0
	total2 := 0
	lines := []Set{}
	n := Set{}

	for scanner.Scan() {

		line := []rune(scanner.Text())

		// Part 1
		n = intersect(toSet(line[:(len(line)/2)]), toSet(line[len(line)/2:]))
		for r := range n {
			total1 = total1 + priority(r)
		}

		// Part 2
		lines = append(lines, toSet(line))
		if (len(lines) % 3) == 0 {
			n = intersect(lines[0], lines[1], lines[2])
			lines = []Set{}
			for r := range n {
				total2 = total2 + priority(r)
			}
		}

	}

	fmt.Printf("Part 1 %v\n", total1)
	fmt.Printf("Part 2 %v\n", total2)
}
