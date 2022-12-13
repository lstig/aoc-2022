package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lstig/aoc-2022/internal/util"
)

type CrateStacks [][]string

func (c CrateStacks) Add(col int, crate string) {
	c[col] = append([]string{crate}, c[col]...)
}

func (c CrateStacks) Move(num int, from int, to int, strategy string) {
	// `from` and `to` columns start from 1
	from--
	to--

	switch strategy {
	case "pop":
		// to store the popped value
		var j string
		for i := 0; i < num; i++ {
			// Pop
			j, c[from] = c[from][len(c[from])-1], c[from][:len(c[from])-1]
			// Push
			c[to] = append(c[to], j)
		}
	case "shift":
		// store slice
		a := c[from][len(c[from])-num:]
		// remove slice from end
		c[from] = c[from][:len(c[from])-num]
		// append slice to new column
		c[to] = append(c[to], a...)
	}

}

func (c CrateStacks) GetTop() string {
	top := []string{}
	for _, col := range c {
		top = append(top, col[len(col)-1])
	}
	return strings.Join(top, "")
}

func main() {
	scanner := util.OpenInput("./input.txt")
	part1 := make(CrateStacks, 9)
	part2 := make(CrateStacks, 9)

	// Get the table from the input
	for scanner.Scan() {
		line := scanner.Text()
		// add the line to the stacks only if it contains crates
		matched, _ := regexp.MatchString(`\s+\[`, line)
		if matched {
			for i, j := 0, 0; i < len(line); i, j = i+4, j+1 {
				crate := string(line[i : i+2][1])
				// ignore blank crates
				if crate != " " {
					part1.Add(j, crate)
					part2.Add(j, crate)
				}
			}
		} else {
			// Keep processing lines until there's a break in the input
			if len(line) == 0 {
				break
			}
		}
	}

	re := regexp.MustCompile(`\d+`)

	// process the moves
	for scanner.Scan() {
		line := scanner.Text()
		move := []int{}

		// parse lines in the format `move x from y to z`
		for _, i := range re.FindAllString(line, -1) {
			j, _ := strconv.Atoi(i)
			move = append(move, j)
		}

		part1.Move(move[0], move[1], move[2], "pop")
		part2.Move(move[0], move[1], move[2], "shift")
	}

	fmt.Printf("Part 1: %v\n", part1.GetTop())
	fmt.Printf("Part 2: %v\n", part2.GetTop())
}
