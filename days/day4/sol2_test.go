package day4

import (
	"bufio"
	"os"
	"testing"
)

func TestRotate(t *testing.T) {
	f, _ := os.Open("test_input2.txt")
	defer f.Close()
	s := bufio.NewScanner(f)
	var elts [][]rune
	for s.Scan() {
		elts = append(elts, []rune(s.Text()))
	}

	g := createGrid(elts)

	if g.Width != 10 {
		t.Errorf("expected width 10, got %d", g.Width)
	}

	if g.Height != 11 {
		t.Errorf("expected width 11, got %d", g.Height)
	}

	g2 := g.CreateRotatedCopy()
	if g2.Width != 11 {
		t.Errorf("expected width 11, got %d", g.Width)
	}

	if g2.Height != 10 {
		t.Errorf("expected width 10, got %d", g.Height)
	}
}

func TestMasXChecking(t *testing.T) {
	f, _ := os.Open("test_input3.txt")
	defer f.Close()
	s := bufio.NewScanner(f)
	var elts [][]rune
	for s.Scan() {
		elts = append(elts, []rune(s.Text()))
	}

	g := createGrid(elts)

	expected := 9
	res := 0
	for i := 0; i < 4; i++ {
		res += findAllMasX(g)
		g = g.CreateRotatedCopy()
	}

	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
