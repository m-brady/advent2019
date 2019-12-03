package day3

import (
	"advent2019/util"
	"fmt"
	"strings"
)


func Part1() {
	lines := util.ReadLines("day3/input.txt")

	wire1 := strings.Split(lines[0],",")
	//wire2 := strings.Split(lines[0],"")


	//var wire1X map[int]int
	//var wire1Y map[int]int
	//var wire2X map[int]int
	//var wire2Y map[int]int


	for _, move := range wire1{
		dir := move[0]
		distance := move[1:]
		switch dir {
		case 'U':
			fmt.Println("hello")
		case 'L':
		case 'D':
		case 'R':
		}
	}



}

