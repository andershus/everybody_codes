package main

import "os"

func main() {
	dat, err := os.ReadFile("part1.in")
	if err != nil {
		panic(err)
	}
	part1(string(dat))
	dat, err = os.ReadFile("part2.in")
	if err != nil {
		panic(err)
	}
	part2(string(dat))
	dat, err = os.ReadFile("part3.in")
	if err != nil {
		panic(err)
	}
	part3(string(dat))
}
