package day3

import (
	"bufio"
	"os"
	"testing"
)

func TestRegex3(t *testing.T) {
	f, _ := os.Open("test_input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	cmds := MulAndCondsRe.FindAllStringSubmatch(scanner.Text(), -1)

	res := processCmds(cmds)
	if res != 18177601 {
		t.Errorf("expected 18177601, got %d", res)
	}
}
