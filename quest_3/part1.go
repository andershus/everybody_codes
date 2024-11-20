package main

import (
	"fmt"
	"strings"
)

var stages = []rune{'#', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':', ';'}

func gprint(grid [][]rune) {

	for _, line := range grid {
		for _, c := range line {
			fmt.Print(string(c))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func valid(i int, j int, snr int, grid [][]rune) bool {
	if snr == 0 {
		return true
	} else if grid[i-1][j] == stages[snr] && grid[i][j-1] == stages[snr] && grid[i+1][j] == stages[snr] && grid[i][j+1] == stages[snr] {
		return true
	}
	return false
}

func duplicate(grid [][]rune) [][]rune {
	duplicate := make([][]rune, len(grid))
	for i := range grid {
		duplicate[i] = make([]rune, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func part1(input string) {
	// 	input = `
	// ..........
	// ..###.##..
	// ...####...
	// ..######..
	// ..######..
	// ...####...
	// ..........`
	var grid [][]rune
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		var l []rune
		for _, c := range line {
			l = append(l, c)
		}
		grid = append(grid, l)
	}
	// gprint(grid)
	count := 0
	for snr, stage := range stages {
		// fmt.Println("Stage:", string(stage))
		new_grid := duplicate(grid)
		for i, line := range grid {
			for j, c := range line {
				if (c == stage) && valid(i, j, snr, grid) {
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
