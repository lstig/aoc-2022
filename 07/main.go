package main

import (
	"fmt"
	"strconv"

	lex "github.com/lstig/aoc-2022/07/lexer"
	"github.com/lstig/aoc-2022/internal/util"
)

type NodeType int

const (
	File NodeType = iota
	Directory
)

type Node struct {
	Name     string
	Type     NodeType
	Children map[string]*Node
	Size     int
	Parent   *Node
}

type FileSystem struct {
	Root       *Node
	WorkingDir *Node
	TotalSize  int
}

func (f *FileSystem) cd(dir string) {
	switch dir {
	case "/":
		f.WorkingDir = f.Root
	case "..":
		f.WorkingDir = f.WorkingDir.Parent
	default:
		d := f.WorkingDir.Children[dir]
		f.WorkingDir = d
	}
}

func (f *FileSystem) mkdir(dir string) {
	f.WorkingDir.Children[dir] = &Node{
		Name:     dir,
		Type:     Directory,
		Parent:   f.WorkingDir,
		Children: map[string]*Node{},
	}
}

func (f *FileSystem) mkfile(file string, size int) {
	f.WorkingDir.Children[file] = &Node{
		Name:   file,
		Type:   File,
		Size:   size,
		Parent: f.WorkingDir,
	}
}

func calculateDirSize(n *Node) int {
	for _, child := range n.Children {
		switch child.Type {
		case File:
			n.Size += child.Size
		case Directory:
			n.Size += calculateDirSize(child)
		}
	}
	return n.Size
}

// Find all directories smaller than 100000
func part1(n *Node) int {
	size := 0
	for _, child := range n.Children {
		if child.Type == Directory {
			if child.Size <= 100000 {
				size += child.Size
			}
			size += part1(child)
		}
	}
	return size
}

// Find the smallest directory to delete that will leave 30000000 available space
func part2(fs *FileSystem) int {
	min := -1
	available := fs.TotalSize - fs.Root.Size

	var traverse func(n *Node)
	traverse = func(n *Node) {
		for _, child := range n.Children {
			if child.Type == Directory {
				if (available + child.Size) >= 30000000 {
					if min == -1 {
						min = child.Size
					} else if child.Size < min {
						min = child.Size
					}
				}
				traverse(child)
			}
		}
	}

	traverse(fs.Root)

	return min
}

func main() {
	s := util.OpenInput("./input.txt")
	lexer := lex.NewLexer(s)
	fs := &FileSystem{
		Root: &Node{
			Name:     "",
			Type:     Directory,
			Children: map[string]*Node{},
		},
		TotalSize: 70000000,
	}

	for tok, str := lexer.Lex(); tok != lex.EOF; tok, str = lexer.Lex() {
		switch tok {
		case lex.DOLLAR:
			continue
		case lex.CD:
			_, dir := lexer.Lex()
			fs.cd(dir)
		case lex.LS:
			continue
		case lex.DIR:
			_, ident := lexer.Lex()
			fs.mkdir(ident)
		case lex.INT:
			size, _ := strconv.Atoi(str)
			_, ident := lexer.Lex()
			fs.mkfile(ident, size)
		}
	}

	fmt.Printf("Total Size: %v\n", calculateDirSize(fs.Root))
	fmt.Printf("Part 1: %v\n", part1(fs.Root))
	fmt.Printf("Part 2: %v\n", part2(fs))
}
