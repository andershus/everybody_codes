package main

import (
	"fmt"
	"os"
)

func part1() {
	dat, err := os.ReadFile("part1.in")
	if err != nil {
		panic(err)
	}
	var count = 0
	for _, c := range string(dat) {
		if c == 'B' {
			count += 1
		}
		if c == 'C' {
			count += 3
		}
	}
	fmt.Println(count)
}
