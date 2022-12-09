package main

import (
	"bufio"
	"fmt"

	"github.com/lstig/aoc-2022/internal/util"
)

type Marker struct {
	// Size of the marker we're looking for
	Size int
	// FIFO queue for holding current set of bytes
	Window []byte
	// Map to hold bytes, if there's a duplicate byte the length of the Map will be less than the Size
	bytes map[byte]struct{}
}

func (m *Marker) Found() bool {
	return len(m.bytes) == m.Size
}

func (m *Marker) Add(b byte) {
	// If the queue is full, pop the first item
	if len(m.Window) == m.Size {
		m.Window = m.Window[1:]
	}
	// Add the new item to the queue and update the map
	m.Window = append(m.Window, b)
	// Reset byte map
	m.bytes = map[byte]struct{}{}
	for _, b := range m.Window {
		m.bytes[b] = struct{}{}
	}

}

func main() {
	scanner := util.OpenInput("./input.txt")
	scanner.Split(bufio.ScanBytes)

	sop, foundP := &Marker{Size: 4, Window: []byte{}}, false
	som, foundM := &Marker{Size: 14, Window: []byte{}}, false

	for i := 0; scanner.Scan(); i++ {
		b := scanner.Bytes()[0]
		sop.Add(b)
		som.Add(b)
		if sop.Found() && !foundP {
			foundP = true
			fmt.Printf("Found packet marker at offset %v: %q\n", i + 1, sop.Window)
		}
		if som.Found() && !foundM {
			foundM = true
			fmt.Printf("Found message marker at offset %v: %q\n", i + 1, som.Window)
		}
	}

}
