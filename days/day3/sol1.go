package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var MulRegex, _ = regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)

func Sol1(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	var factors [][2]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		newFactors := extractFactors(line)
		factors = append(factors, newFactors...)
	}

	res := 0
	for _, pair := range factors {
		res += (pair[0] * pair[1])
	}

	return fmt.Sprintf("%d", res)
}

func extractFactors(line string) (res [][2]int) {
	matches := MulRegex.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		xS, yS := match[len(match)-2], match[len(match)-1]
		x, _ := strconv.Atoi(xS)
		y, _ := strconv.Atoi(yS)
		newFactors := [2]int{x, y}
		res = append(res, newFactors)
	}
	return
}
