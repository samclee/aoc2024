package day4

import (
	"bufio"
	"os"
	"testing"
)

func TestCreateGrid(t *testing.T) {
	f, _ := os.Open("test_input.txt")
	defer f.Close()
	s := bufio.NewScanner(f)
	var elts [][]rune
	for s.Scan() {
		elts = append(elts, []rune(s.Text()))
	}

	g := createGrid(elts)

	if g.Width != 7 {
		t.Errorf("expected width 7, got %d", g.Width)
	}

	if g.Height != 5 {
		t.Errorf("expected width 5, got %d", g.Height)
	}
}
