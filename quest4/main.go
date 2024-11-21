package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func parse(input string) []int {
	input = `
3
4
7
8
`
	var input_p []int
	for _, d := range strings.Split(strings.Trim(input, "\n"), "\n") {
		i, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		input_p = append(input_p, i)
	}
	return input_p
}

func part1(input string) {
	input_p := parse(input)
	m := slices.Min(input_p)
	count := 0
	for _, d := range input_p {
		count += d - m
	}
	fmt.Println(count)

}

func part3(input string) {
	input = `
2
4
5
6
8
`
	input_p := parse(input)
	sort.Sort(sort.IntSlice(input_p))
	median := input_p[len(input_p)/2]
	count := 0
	for _, d := range input_p {
		diff := d - median
		if diff < 0 {
			diff = -diff
		}
		count += diff
	}
	fmt.Println(count)
}

func main() {
	dat, err := os.ReadFile("part1.in")
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
		part1(string(dat))
	}
	dat, err = os.ReadFile("part3.in")
	if err != nil {
		panic(err)
	}
	if len(dat) != 0 {
		part3(string(dat))
	}
}
