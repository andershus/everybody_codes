package main

import (
	"fmt"
	"strings"
)

func valid_diag(i int, j int, snr int, grid [][]rune) bool {
	if snr == 0 {
		return true
	} else if grid[i-1][j-1] == stages[snr] && grid[i-1][j] == stages[snr] && grid[i-1][j+1] == stages[snr] && grid[i][j-1] == stages[snr] && grid[i][j+1] == stages[snr] && grid[i+1][j-1] == stages[snr] && grid[i+1][j] == stages[snr] && grid[i+1][j+1] == stages[snr] {
		return true
	}
	return false
}

func part3(input string) {
	// 	input = `
	// ###.##
	// .####.
	// ######
	// ######
	// .####.`
	input_list := strings.Split(strings.Trim(input, "\n"), "\n")

	var grid [][]rune
	empty_line := make([]rune, len(input_list[0])+2)
	for i := range empty_line {
		empty_line[i] = '.'
	}
	grid = append(grid, empty_line)
	for _, line := range input_list {
		var l []rune
		l = append(l, '.')
		for _, c := range line {
			l = append(l, c)
		}
		l = append(l, '.')
		grid = append(grid, l)
	}
	grid = append(grid, empty_line)
	// gprint(grid)

	count := 0
	for snr, stage := range stages {
		fmt.Println("Stage:", string(stage))
		new_grid := duplicate(grid)
		for i, line := range grid {
			for j, c := range line {
				if (c == stage) && valid_diag(i, j, snr, grid) {
					count++
					new_grid[i][j] = stages[snr+1]
				}
			}
		}
		grid = new_grid
		// gprint(grid)
		// fmt.Println(count)
	}
	fmt.Println(count)
}
