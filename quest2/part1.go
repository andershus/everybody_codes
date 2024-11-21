package main

import (
	"fmt"
	"strings"
)

func part1(input string) {
	// 	input = `WORDS:THE,OWE,MES,ROD,HER
	//
	// AWAKEN THE POWER ADORNED WITH THE FLAMES BRIGHT IRE`
	input_p := strings.Split(input, "\n\n")
	words := strings.Split(strings.Split(input_p[0], ":")[1], ",")
	sentence := input_p[1]
	var count = 0
	for _, word := range words {
		count += strings.Count(sentence, word)

	}
	fmt.Println(count)
}
