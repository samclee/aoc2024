package cmd

import (
	"aoc2024/days/day1"
	"aoc2024/days/day10"
	"aoc2024/days/day11"
	"aoc2024/days/day12"
	"aoc2024/days/day13"
	"aoc2024/days/day14"
	"aoc2024/days/day15"
	"aoc2024/days/day16"
	"aoc2024/days/day17"
	"aoc2024/days/day18"
	"aoc2024/days/day19"
	"aoc2024/days/day2"
	"aoc2024/days/day20"
	"aoc2024/days/day21"
	"aoc2024/days/day22"
	"aoc2024/days/day23"
	"aoc2024/days/day24"
	"aoc2024/days/day25"
	"aoc2024/days/day3"
	"aoc2024/days/day4"
	"aoc2024/days/day5"
	"aoc2024/days/day6"
	"aoc2024/days/day7"
	"aoc2024/days/day8"
	"aoc2024/days/day9"
	"fmt"
)

func Solve(day int, part int, inputFname string) (string, error) {
	switch day {
	case 1:
		if part == 1 {
			return day1.Sol1(inputFname), nil
		} else {
			return day1.Sol2(inputFname), nil
		}
	case 2:
		if part == 1 {
			return day2.Sol1(inputFname), nil
		} else {
			return day2.Sol2(inputFname), nil
		}
	case 3:
		if part == 1 {
			return day3.Sol1(inputFname), nil
		} else {
			return day3.Sol2(inputFname), nil
		}
	case 4:
		if part == 1 {
			return day4.Sol1(inputFname), nil
		} else {
			return day4.Sol2(inputFname), nil
		}
	case 5:
		if part == 1 {
			return day5.Sol1(inputFname), nil
		} else {
			return day5.Sol2(inputFname), nil
		}
	case 6:
		if part == 1 {
			return day6.Sol1(inputFname), nil
		} else {
			return day6.Sol2(inputFname), nil
		}
	case 7:
		if part == 1 {
			return day7.Sol1(inputFname), nil
		} else {
			return day7.Sol2(inputFname), nil
		}
	case 8:
		if part == 1 {
			return day8.Sol1(inputFname), nil
		} else {
			return day8.Sol2(inputFname), nil
		}
	case 9:
		if part == 1 {
			return day9.Sol1(inputFname), nil
		} else {
			return day9.Sol2(inputFname), nil
		}
	case 10:
		if part == 1 {
			return day10.Sol1(inputFname), nil
		} else {
			return day10.Sol2(inputFname), nil
		}
	case 11:
		if part == 1 {
			return day11.Sol1(inputFname), nil
		} else {
			return day11.Sol2(inputFname), nil
		}
	case 12:
		if part == 1 {
			return day12.Sol1(inputFname), nil
		} else {
			return day12.Sol2(inputFname), nil
		}
	case 13:
		if part == 1 {
			return day13.Sol1(inputFname), nil
		} else {
			return day13.Sol2(inputFname), nil
		}
	case 14:
		if part == 1 {
			return day14.Sol1(inputFname), nil
		} else {
			return day14.Sol2(inputFname), nil
		}
	case 15:
		if part == 1 {
			return day15.Sol1(inputFname), nil
		} else {
			return day15.Sol2(inputFname), nil
		}
	case 16:
		if part == 1 {
			return day16.Sol1(inputFname), nil
		} else {
			return day16.Sol2(inputFname), nil
		}
	case 17:
		if part == 1 {
			return day17.Sol1(inputFname), nil
		} else {
			return day17.Sol2(inputFname), nil
		}
	case 18:
		if part == 1 {
			return day18.Sol1(inputFname), nil
		} else {
			return day18.Sol2(inputFname), nil
		}
	case 19:
		if part == 1 {
			return day19.Sol1(inputFname), nil
		} else {
			return day19.Sol2(inputFname), nil
		}
	case 20:
		if part == 1 {
			return day20.Sol1(inputFname), nil
		} else {
			return day20.Sol2(inputFname), nil
		}
	case 21:
		if part == 1 {
			return day21.Sol1(inputFname), nil
		} else {
			return day21.Sol2(inputFname), nil
		}
	case 22:
		if part == 1 {
			return day22.Sol1(inputFname), nil
		} else {
			return day22.Sol2(inputFname), nil
		}
	case 23:
		if part == 1 {
			return day23.Sol1(inputFname), nil
		} else {
			return day23.Sol2(inputFname), nil
		}
	case 24:
		if part == 1 {
			return day24.Sol1(inputFname), nil
		} else {
			return day24.Sol2(inputFname), nil
		}
	case 25:
		if part == 1 {
			return day25.Sol1(inputFname), nil
		} else {
			return day25.Sol2(inputFname), nil
		}
	default:
		return "", fmt.Errorf("day [%d] not found", day)
	}
}
