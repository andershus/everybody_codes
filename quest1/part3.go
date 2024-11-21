package main

import (
	"fmt"
	"os"
)

func part3() {
	dat, err := os.ReadFile("part3.in")
	if err != nil {
		panic(err)
	}
	str := string(dat)
	// str = "xBxAAABCDxCC"
	count := 0
	for i := 0; i < len(str)-2; i += 3 {
		pair := str[i : i+3]
		extra := 2
		for _, c := range pair {
			if c == 'x' {
				extra--
			}
			if c == 'A' {
				count += 0
			}
			if c == 'B' {
				count += 1
			}
			if c == 'C' {
				count += 3
			}
			if c == 'D' {
				count += 5
			}
		}
		count += extra * (extra + 1)
	}
	fmt.Println(count)
}
