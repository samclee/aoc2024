package day16

import "os"

func Sol1(fname string) string {
	f, _ := os.Open(fname)
	defer f.Close()

	return ""
}