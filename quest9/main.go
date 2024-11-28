package main

import (
	"fmt"
	"os"
	// "slices"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func parse(input string) []int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	parsed := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		parsed[i] = n
	}
	return parsed
}

func part1(input string) {
	// 	input = `
	// 2
	// 4
	// 7
	// 16
	// `
	brightnesses := parse(input)
	stamps := []int{10, 5, 3, 1}
	count := 0
	for _, brightness := range brightnesses {
		for _, stamp := range stamps {
			for stamp <= brightness {
				count++
				brightness -= stamp
			}
		}
	}
	fmt.Println("Part1:", count)
}

// Minimum change coin problem: https://stackoverflow.com/questions/9964289/finding-shortest-combinations-in-array-sequence-that-equals-sum
func min_change(V []int, C int) []int {
	table, solution := min_change_table(V, C)
	num_coins, coins := table[len(table)-1], []int{}
	if num_coins == MaxInt {
		return []int{}
	}
	for C > 0 {
		coins = append(coins, V[solution[C]])
		C -= V[solution[C]]
	}
	return coins
}

func min_change_table(V []int, C int) ([]int, []int) {
	m, n := C+1, len(V)
	var table, solution []int
	for range m {
		table = append(table, 0)
		solution = append(solution, 0)
	}
	for i := range m - 1 {
		minNum, minIdx := MaxInt, -1
		for j := range n {
			if V[j] <= i+1 && 1+table[i+1-V[j]] < minNum {
				minNum = 1 + table[i+1-V[j]]
				minIdx = j
			}
		}
		table[i+1] = minNum
		solution[i+1] = minIdx
	}
	return table, solution
}

func part2(input string) {
	// 	input = `
	// 33
	// 41
	// 55
	// 99
	// `
	brightnesses := parse(input)
	stamps := []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30}
	count := 0
	for _, brightness := range brightnesses {
		count += len(min_change(stamps, brightness))
	}
	fmt.Println("Part2:", count)
}

func part3(input string) {
	// 	input = `
	// 156488
	// 352486
	// 546212
	// `
	brightnesses := parse(input)
	stamps := []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}
	count := 0
	for i, brightness := range brightnesses {
		fmt.Println("Processing", i, "of", len(brightnesses)) // this takes about 0.5 seconds to run each loop
		offset := 0
		if brightness%2 == 1 {
			offset = 1
		}
		mid := brightness / 2
		mini := MaxInt
		for i := range 51 {
			if 2*i+offset > 100 {
				break
			}
			stamp1 := len(min_change(stamps, mid-i))
			stamp2 := len(min_change(stamps, mid+i+offset))
			if stamp1+stamp2 < mini {
				mini = stamp1 + stamp2
			}

		}
		count += mini
	}
	fmt.Println("Part3:", count)
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
