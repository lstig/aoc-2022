package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Heap struct {
	Heap []int
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) rightChild(i int) int {
	return  (2 * i) + 2
}

func (h *Heap) leftChild(i int) int {
	return (2 * i) + 1
}

func (h *Heap) isLeaf(i int) bool {
	return i > (len(h.Heap) / 2) && i <= len(h.Heap)
}

func (h *Heap) swap(i int, j int) {
	tmp := h.Heap[i]
	h.Heap[i] = h.Heap[j]
	h.Heap[j] = tmp
}

func (h *Heap) makeHeap(i int) {
	if h.isLeaf(i) {
		return
	}

	c, l, r := h.Heap[i], h.Heap[h.leftChild(i)], h.Heap[h.rightChild(i)]

	if (c < l) || (c < r) {
		if l > r {
			h.swap(i, h.leftChild(i))
			h.makeHeap(h.leftChild(i))
		} else {
			h.swap(i, h.rightChild(i))
			h.makeHeap(h.rightChild(i))
		}
	}
}

func (h *Heap) Pop() int {
	max := h.Heap[0]
	h.Heap = h.Heap[1:]
	h.makeHeap(0)
	return max
}

func (h *Heap) Insert(i int) {
	current := len(h.Heap)
	h.Heap = append(h.Heap, i)
	for h.Heap[current] > h.Heap[h.parent(current)] {
		h.swap(current, h.parent(current))
		current = h.parent(current)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heap := Heap{ Heap: []int{} }
	current := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			current = current + calories
			continue
		}
		heap.Insert(current)
		current = 0
	}

	fmt.Printf("Most Calories %v\n", heap.Heap[0])
	fmt.Printf("Top 3 %v\n", heap.Pop() + heap.Pop() + heap.Pop())
}