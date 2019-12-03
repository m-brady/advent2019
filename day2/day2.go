package day2

import (
	"advent2019/util"
	"fmt"
	"strconv"
	"strings"
)

func Part1() {

	lines := util.ReadLines("day2/input.txt")

	line := strings.Split(lines[0], ",")
	codes := make([]int, len(line))
	for i, c := range line {
		codes[i], _ = strconv.Atoi(c)
	}
	codes[1] = 12
	codes[2] = 2

	for i := 0; i < len(codes); i += 4 {
		c := codes[i]
		if c == 1 {
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
		} else if c == 2 {
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
		} else if c == 99 {
			fmt.Println(codes[0])
		}

	}
}

func Part2() {
	output := 19690720
	lines := util.ReadLines("day2/input.txt")

	line := strings.Split(lines[0], ",")

	for noun := 0; noun <= 99; noun++ {

		for verb := 0; verb <= 99; verb++ {

			codes := make([]int, len(line))
			for i, c := range line {
				codes[i], _ = strconv.Atoi(c)
			}
			codes[1] = noun
			codes[2] = verb

			for i := 0; i < len(codes); i += 4 {
				c := codes[i]
				if c == 1 {
					codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
				} else if c == 2 {
					codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
				} else if c == 99 {
					if codes[0] == output {
						fmt.Println(100 * noun + verb)
					}
 					break
				}
			}
		}

	}

}
