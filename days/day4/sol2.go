package day4

import (
	"bufio"
	"fmt"
	"os"
)

func (g Grid) CreateRotatedCopy() Grid {
	newG := Grid{
		Width:  g.Height,
		Height: g.Width,
	}

	newElts := make([][]rune, g.Width)
	for i := range newElts {
		newElts[i] = make([]rune, g.Height)
	}
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			newElts[c][r] = g.elements[r][g.Width-1-c]
		}
	}
	newG.elements = newElts

	return newG
}

func Sol2(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	var elts [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elts = append(elts, []rune(scanner.Text()))
	}
	g := createGrid(elts)

	res := 0
	for i := 0; i < 4; i++ {
		res += findAllMasX(g)
		g = g.CreateRotatedCopy()
	}

	return fmt.Sprintf("%d", res)
}

func findAllMasX(g Grid) int {
	res := 0
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			res += isMasX(r, c, g)
		}
	}

	return res
}

func isMasX(row, col int, g Grid) int {
	if center, ok := g.GetElement(row, col); !ok || center != 'A' {
		return 0
	}

	tl, ok1 := g.GetElement(row-1, col-1)
	tr, ok2 := g.GetElement(row-1, col+1)
	bl, ok3 := g.GetElement(row+1, col-1)
	br, ok4 := g.GetElement(row+1, col+1)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		return 0
	}
	if tl != 'S' || tr != 'S' || bl != 'M' || br != 'M' {
		return 0
	}

	return 1
}
