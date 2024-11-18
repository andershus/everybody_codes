package main

import (
	"fmt"
	"os"
)

func part2() {
	dat, err := os.ReadFile("part2.in")
	if err != nil {
		panic(err)
	}
	str := string(dat)
	// str = "AxBCDDCAxD"
	count := 0
	for i := 0; i < len(str)-1; i += 2 {
		extra := 0
		pair := str[i : i+2]
		if pair[0] != 'x' && pair[1] != 'x' {
			extra = 1
		}
		for _, c := range pair {
			if c == 'A' {
				count += 0 + extra
			}
			if c == 'B' {
				count += 1 + extra
			}
			if c == 'C' {
				count += 3 + extra
			}
			if c == 'D' {
				count += 5 + extra
			}
		}
	}
	fmt.Println(count)
}
