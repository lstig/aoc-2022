package main

import (
	"fmt"

	"github.com/lstig/aoc-2022/internal/util"
)

// Rock     A X - 1pt
// Paper    B Y - 2pt
// Scissors C Z - 3pt
// Lose 0pt
// Draw 3pt
// Win  6pt

var ScoreTablePart1 = map[string]int{
	"A X": 4, // 1 + 3
	"A Y": 8, // 2 + 6
	"A Z": 3, // 3 + 0
	"B X": 1, // 1 + 0
	"B Y": 5, // 2 + 3
	"B Z": 9, // 3 + 6
	"C X": 7, // 1 + 6
	"C Y": 2, // 2 + 0
	"C Z": 6, // 3 + 3
}

var ScoreTablePart2 = map[string]int{
	"A X": 3, // 3 + 0
	"A Y": 4, // 1 + 3
	"A Z": 8, // 2 + 6
	"B X": 1, // 1 + 0
	"B Y": 5, // 2 + 3
	"B Z": 9, // 3 + 6
	"C X": 2, // 2 + 0
	"C Y": 6, // 3 + 3
	"C Z": 7, // 1 + 6
}

func main() {
	scanner := util.OpenInput("./input.txt")
	total1 := 0
	total2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		total1 = total1 + ScoreTablePart1[line]
		total2 = total2 + ScoreTablePart2[line]
	}

	fmt.Printf("Part 1 Score: %v\n", total1)
	fmt.Printf("Part 2 Scare: %v\n", total2)
}
