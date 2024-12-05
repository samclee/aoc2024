package day4

import (
	"bufio"
	"fmt"
	"os"
)

type Grid struct {
	elements [][]rune
	Width    int
	Height   int
}

func createGrid(elements [][]rune) Grid {
	g := Grid{
		elements: elements,
		Width:    len(elements[0]),
		Height:   len(elements),
	}

	return g
}

func (g Grid) GetElement(row, col int) (rune, bool) {
	if g.IsOob(row, col) {
		return 0, false
	}

	return g.elements[row][col], true
}

func (g Grid) IsOob(row, col int) bool {
	if row < 0 || col < 0 || row > g.Height-1 || col > g.Width-1 {
		return true
	}

	return false
}

func (g Grid) Print() {
	for _, r := range g.elements {
		for _, c := range r {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

var dirs = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

var order = [4]rune{
	'X',
	'M',
	'A',
	'S',
}

func Sol1(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	var elements [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elements = append(elements, []rune(scanner.Text()))
	}
	g := createGrid(elements)

	insts := 0
	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			for _, d := range dirs {
				insts += xmasExistsInDir(row, col, d[0], d[1], g)
			}
		}
	}

	return fmt.Sprintf("%d", insts)
}

func xmasExistsInDir(row int, col int, dx int, dy int, grid Grid) int {
	for i := 0; i < 4; i++ {
		nextRune, ok := grid.GetElement(row+i*dx, col+i*dy)
		if !ok {
			return 0
		}
		if nextRune != order[i] {
			return 0
		}
	}

	return 1
}
