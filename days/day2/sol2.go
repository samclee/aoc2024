package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Sol2(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	res := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		ints := toIntSlice(tokens)
		if diffsInRange(ints, 1, 3) || diffsInRange(ints, -3, -1) || diffsInRangeWithIgnore(ints, 1, 3) || diffsInRangeWithIgnore(ints, -3, -1) {
			res += 1
		}
	}

	return fmt.Sprintf("%d", res)
}

func diffsInRangeWithIgnore(slice []int, low, high int) bool {
	for i := 0; i < len(slice); i++ {
		newSlice := removeElt(slice, i)
		if diffsInRange(newSlice, low, high) {
			return true
		}
	}

	return false
}

func removeElt(slice []int, i int) []int {
	var newSlice []int
	newSlice = append(newSlice, slice[:i]...)
	newSlice = append(newSlice, slice[i+1:]...)

	return newSlice
}
