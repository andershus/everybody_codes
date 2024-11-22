package main

import (
	"fmt"
	"os"
	"slices"
	// "sort"
	"strconv"
	"strings"
)

func parse(input string) [][]int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	parsed := make([][]int, len(strings.Split(lines[0], " ")))
	for _, l := range lines {
		for j, c := range strings.Split(l, " ") {
			d, _ := strconv.Atoi(c)
			parsed[j] = append(parsed[j], d)
		}
	}
	return parsed
}

func part1(input string) {
	// 	input = `
	// 2 3 4 5
	// 1 4 5 2
	// 4 5 2 3
	// 5 2 3 4
	// `
	columns := parse(input)
	num_cols := len(columns)
	for round := range 10 {
		cur_col := round % num_cols
		next_col := (cur_col + 1) % num_cols
		var clapper int
		clapper, columns[cur_col] = columns[cur_col][0], columns[cur_col][1:]
		// fmt.Println("Round", round+1, "Clapper:", clapper)
		len_ncol := len(columns[next_col])
		if clapper <= len_ncol {
			columns[next_col] = slices.Insert(columns[next_col], clapper-1, clapper)

		} else {
			a := clapper - len_ncol - 1 // golang wierd behaviour
			columns[next_col] = slices.Insert(columns[next_col], a+1, clapper)
			fmt.Println(columns)
		}
		// fmt.Println(cur_col, next_col, clapper, columns)
	}
	result := ""
	for _, col := range columns {
		result += strconv.Itoa(col[0])
	}
	// fmt.Println(result)
}

func part2(input string) {
	// 	input = `
	// 2 3 4 5
	// 6 7 8 9
	// `
	columns := parse(input)
	// fmt.Println(columns)
	num_cols := len(columns)
	shouts := map[string]int{}
	for round := 0; round < 10_000_000; round++ {
		cur_col := round % num_cols
		next_col := (cur_col + 1) % num_cols
		var clapper int
		clapper, columns[cur_col] = columns[cur_col][0], columns[cur_col][1:]
		// fmt.Println("Round", round+1, "Clapper:", clapper)
		len_ncol := len(columns[next_col])
		new_pos := (clapper - 1) % (2 * len_ncol)
		// fmt.Println("We go to pass", pass)
		if new_pos >= len_ncol {
			new_pos = 2*len_ncol - new_pos
		}
		columns[next_col] = slices.Insert(columns[next_col], new_pos, clapper)
		// fmt.Println(columns)
		result := ""
		for _, col := range columns {
			result += strconv.Itoa(col[0])
		}
		shouts[result]++
		// fmt.Println(result, shouts[result])
		if shouts[result] == 2024 {
			r, err := strconv.Atoi(result)
			if err != nil {
				panic(err)
			}
			fmt.Println(round+1, r, (round+1)*r) // wrong len, wrong first char: 809456756031
			break
		}
	}
}

func part3(input string) {
	// 	input = `
	// 2 3 4 5
	// 6 7 8 9
	// `
	columns := parse(input)
	num_cols := len(columns)
	highest := 0
	seen := map[string]struct{}{}
	for round := 0; round >= 0; round++ {
		cur_col := round % num_cols
		next_col := (cur_col + 1) % num_cols
		var clapper int
		clapper, columns[cur_col] = columns[cur_col][0], columns[cur_col][1:]
		// fmt.Println("Round", round+1, "Clapper:", clapper)
		len_ncol := len(columns[next_col])
		new_pos := (clapper - 1) % (2 * len_ncol)
		// fmt.Println("We go to pass", pass)
		if new_pos >= len_ncol {
			new_pos = 2*len_ncol - new_pos
		}
		columns[next_col] = slices.Insert(columns[next_col], new_pos, clapper)
		// fmt.Println(columns)
		result := ""
		for _, col := range columns {
			result += strconv.Itoa(col[0])
		}
		r, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		if r > highest {
			highest = r
			fmt.Println(highest)
		}
		state := ""
		for _, col := range columns {
			for _, p := range col {
				state += strconv.Itoa(p) + " "
			}
			state += "\n"
		}
		if _, ok := seen[state]; ok {
			break
		}
		seen[state] = struct{}{}
	}
}

func main() {
	// dat, err := os.ReadFile("part1.in")
	// if err != nil {
	// 	panic(err)
	// }
	// if len(dat) != 0 {
	// 	part1(string(dat))
	// }
	// dat, err := os.ReadFile("part2.in")
	// if err != nil {
	// 	panic(err)
	// }
	// if len(dat) != 0 {
	// 	part2(string(dat))
	// }
	dat, err := os.ReadFile("part3.in")
	if err != nil {
		panic(err)
	}
	if len(dat) != 0 {
		part3(string(dat))
	}
}
