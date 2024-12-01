package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sol2(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()
	var slice []int
	appearanceMap := make(map[int]int)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		tokens := strings.Split(line, "   ")
		item1, _ := strconv.Atoi(tokens[0])
		slice = append(slice, item1)
		item2, _ := strconv.Atoi(tokens[1])

		if val, ok := appearanceMap[item2]; ok {
			appearanceMap[item2] = val + 1
		} else {
			appearanceMap[item2] = 1
		}
	}

	res := 0
	for _, elt := range slice {
		if appearanceFactor, ok := appearanceMap[elt]; ok {
			res += elt * appearanceFactor
		}
	}

	return fmt.Sprintf("%d", res)
}
