package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/lstig/aoc-2022/internal/util"
)

func ranges(s string) []int {
	ranges := []int{}
	for _, i := range strings.Split(s, ",") {
		for _, j := range strings.Split(i, "-") {
			v, _ := strconv.Atoi(j)
			ranges = append(ranges, v)
		}
	}
	return ranges
}

func contains(l1 int, r1 int, l2 int, r2 int) int {
	if (l1 >= l2 && r1 <= r2) || (l2 >= l1 && r2 <= r1) {
		return 1
	}
	return 0
}

func overlap(l1 int, r1 int, l2 int, r2 int) int {
	if math.Max(float64(l1), float64(l2)) <= math.Min(float64(r1), float64(r2)) {
		return 1
	}
	return 0
}

func main() {
	scanner := util.OpenInput("./input.txt")
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		line := ranges(scanner.Text())
		part1 = part1 + contains(line[0], line[1], line[2], line[3])
		part2 = part2 + overlap(line[0], line[1], line[2], line[3])

	}

	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}
