package main

import (
	"fmt"
	"os"
	// "slices"
	// "sort"
	// "strconv"
	"strings"
)

func parse(input string) map[string][]string {
	parsed := make(map[string][]string)
	for _, line := range strings.Split(strings.Trim(input, "\n"), "\n") {
		parts := strings.Split(line, ":")
		parsed[parts[0]] = strings.Split(parts[1], ",")
	}
	return parsed
}

func find_paths(graph map[string][]string, path []string) [][]string {
	var paths [][]string
	start := path[len(path)-1]
	if start == "@" {
		return [][]string{{"@"}}
	} else if start == "BUG" || start == "ANT" { // only for part 3
		return [][]string{{""}}
	}
	for _, e := range graph[start] {
		for _, npath := range find_paths(graph, append(path, e)) {
			paths = append(paths, append([]string{start}, npath...))
		}

	}
	return paths
}

func part1(input string) {
	// 	input = `
	// RR:A,B,C
	// A:D,E
	// B:F,@
	// C:G,H
	// D:@
	// E:@
	// F:@
	// G:@
	// H:@
	// `
	graph := parse(input)
	// fmt.Println(graph)
	paths := find_paths(graph, []string{"RR"})
	// fmt.Println(paths)
	lengths := make(map[int][][]string)
	for _, path := range paths {
		lengths[len(path)] = append(lengths[len(path)], path)
	}
	var unique []string
	for _, length := range lengths {
		if len(length) == 1 {
			unique = length[0]
		}
	}
	fmt.Println(strings.Join(unique, ""))
}

func part2(input string) {
	graph := parse(input)
	paths := find_paths(graph, []string{"RR"})
	lengths := make(map[int][][]string)
	for _, path := range paths {
		lengths[len(path)] = append(lengths[len(path)], path)
	}
	var unique []string
	for _, length := range lengths {
		if len(length) == 1 {
			unique = length[0]
		}
	}
	firstc := ""
	for _, e := range unique {
		firstc += string(e[0])
	}
	fmt.Println(firstc)
}

func part3(input string) {
	graph := parse(input)
	paths := find_paths(graph, []string{"RR"})
	lengths := make(map[int][][]string)
	for _, path := range paths {
		lengths[len(path)] = append(lengths[len(path)], path)
	}
	var unique []string
	for _, length := range lengths {
		if len(length) == 1 {
			unique = length[0]
		}
	}
	firstc := ""
	for _, e := range unique {
		firstc += string(e[0])
	}
	fmt.Println(firstc)
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
