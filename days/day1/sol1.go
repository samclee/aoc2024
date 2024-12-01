package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Sol1(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()
	var slice1 []int
	var slice2 []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		tokens := strings.Split(line, "   ")
		item1, _ := strconv.Atoi(tokens[0])
		slice1 = append(slice1, item1)
		item2, _ := strconv.Atoi(tokens[1])
		slice2 = append(slice2, item2)
	}

	slices.Sort(slice1)
	slices.Sort(slice2)

	res := 0
	for i, elt1 := range slice1 {
		elt2 := slice2[i]
		res += absDiff(elt1, elt2)
	}

	return fmt.Sprintf("%d", res)
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
