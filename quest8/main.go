package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(input string) {
	// input = `13`
	blocks, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		panic(err)
	}
	total := 0
	width := 1
	for total < blocks {
		total += width
		width += 2
	}
	fmt.Println("Part1:", total, width-2, (total-blocks)*(width-2))
}

func part2(input string) {
	acolytes := 1111
	blocks := 20240000
	// input = `3` // test input
	// acolytes := 5 // test input
	// blocks := 50 // test input
	notes, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		panic(err)
	}
	total := 1
	width := 1
	thickness := 1
	for total < blocks {
		thickness = (thickness * notes) % acolytes
		width += 2
		total += width * thickness
	}
	fmt.Println("Part2:", total, width, (total-blocks)*(width))
}

func part3(input string) {
	acolytes := 10
	blocks := 202400000
	// input = `2`  // test input
	// acolytes = 5 // test input
	// blocks = 160 // test input

	notes, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		panic(err)
	}
	total := 1
	width := 1
	thickness := 1
	heights := []int{1}
	for total < blocks {
		thickness = (thickness*notes)%acolytes + acolytes
		for i := range heights {
			heights[i] += thickness
		}
		heights = append(heights, thickness)
		width += 2
		total += width * thickness
	}
	for i, height := range heights {
		to_remove := (notes * width * height) % acolytes
		if i == 0 {
			total -= to_remove
		} else {
			if i != len(heights)-1 {
				total -= 2 * to_remove
			}
		}
	}
	fmt.Println("Part3:", total, (total - blocks))
}

func main() {
	var dat []byte
	var err error
	dat, err = os.ReadFile("part1.in")
	if err != nil {
		panic(err)
	}
	if len(dat) != 0 {
		part1(string(dat))
	}
	dat, err = os.ReadFile("part2.in")
	if err != nil {
		panic(err)
	}
	if len(dat) != 0 {
		part2(string(dat))
	}
	dat, err = os.ReadFile("part3.in")
	if err != nil {
		panic(err)
	}
	if len(dat) != 0 {
		part3(string(dat))
	}
}
