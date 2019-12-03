package day1

import (
	"advent2019/util"
	"fmt"
	"strconv"
)

func Part1() {

	lines := util.ReadLines("day1/input.txt")

	totalFuel := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		totalFuel += (mass / 3) - 2
	}

	fmt.Println(totalFuel)

}

func Part2() {

	lines := util.ReadLines("day1/input.txt")

	totalFuel := 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		fuel := (mass / 3) - 2
		for ; fuel > 0; {
			totalFuel += fuel
			fuel = (fuel / 3) - 2
		}

	}

	fmt.Println(totalFuel)

}


