package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	DO   = "do()"
	DONT = "don't()"
)

var MulAndCondsRe, _ = regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func Sol2(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	var cmds [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		newCmds := MulAndCondsRe.FindAllStringSubmatch(line, -1)
		cmds = append(cmds, newCmds...)
	}

	res := processCmds(cmds)

	return fmt.Sprintf("%d", res)
}

func processCmds(cmds [][]string) int {
	enabledFactor := 1
	res := 0
	for _, c := range cmds {
		if !strings.Contains(c[0], "mul") {
			if c[0] == DO {
				enabledFactor = 1
			} else if c[0] == DONT {
				enabledFactor = 0
			}
		} else {
			xS, yS := c[len(c)-2], c[len(c)-1]
			x, _ := strconv.Atoi(xS)
			y, _ := strconv.Atoi(yS)
			res += (x * y * enabledFactor)
		}
	}

	return res
}
