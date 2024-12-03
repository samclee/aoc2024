package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sol1(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	res := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		ints := toIntSlice(tokens)
		if diffsInRange(ints, 1, 3) || diffsInRange(ints, -3, -1) {
			res += 1
		}
	}

	return fmt.Sprintf("%d", res)
}

func diffsInRange(slice []int, low, high int) bool {
	for i, elt := range slice {
		if i == 0 {
			continue
		}
		diff := elt - slice[i-1]
		if !(low <= diff && diff <= high) {
			return false
		}
	}

	return true
}

func toIntSlice(strSlice []string) (res []int) {
	for _, s := range strSlice {
		i, _ := strconv.Atoi(s)
		res = append(res, i)
	}

	return
}
