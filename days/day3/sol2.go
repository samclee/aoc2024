package day3

import "os"

func Sol2(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	return ""
}