package main

import (
	"fmt"
	"os"
	"sort"
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

type kv struct {
	Key   string
	Value int
}

func part1(input string) {
	// 	input = `
	// A:+,-,=,=
	// B:+,=,-,+
	// C:=,-,+,+
	// D:=,=,=,+
	// `
	plans := parse(input)
	num_rounds := 10
	output := make(map[string][]int)
	results := make(map[string]int)
	for k, plan := range plans {
		power := 10
		for round := range num_rounds {
			len_plan := len(plan)
			index := round % len_plan
			switch plan[index] {
			case "+":
				power++
			case "-":
				if power > 0 {
					power--
				}
			default:
			}
			output[k] = append(output[k], power)
		}
		for _, v := range output[k] {
			results[k] += v
		}
	}
	// fmt.Println(output)
	// fmt.Println(results)
	var ss []kv
	for k, v := range results {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s", kv.Key)
	}
	fmt.Printf("\n")
}

func parse_grid(input string) [][]string {
	var grid [][]string
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	len_line := len(lines[0]) // We assume the first line is complete
	for _, l := range lines {
		var line []string
		for _, c := range strings.Split(l, "") {
			line = append(line, c)
		}
		if len(line) < len_line {
			for range len_line - len(line) {
				line = append(line, " ")
			}
		}
		grid = append(grid, line)
	}
	return grid
}

func parse_racetrack(input string) []string {
	grid := parse_grid(input)
	n := len(grid)
	m := len(grid[0])
	var parsed []string
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	di := 0
	pos := []int{0, 1}
	for pos[0] != 0 || pos[1] != 0 {
		cur_c := grid[pos[0]][pos[1]]
		parsed = append(parsed, cur_c)
		cur_dir := directions[di]
		new_i := pos[0] + cur_dir[0]
		new_j := pos[1] + cur_dir[1]
		if new_i < 0 || n <= new_i || new_j < 0 || m <= new_j {
			di = (di + 1) % 4
			cur_dir = directions[di]
		}
		pos[0] = pos[0] + cur_dir[0]
		pos[1] = pos[1] + cur_dir[1]
	}
	parsed = append(parsed, grid[0][0])
	return parsed
}

func getResult(parsed_racetrack []string, plan []string, loops int) int {
	result := 0
	power := 10
	segment := 0
	len_plan := len(plan)
	for range loops {
		for _, track := range parsed_racetrack {
			index := segment % len_plan
			switch track {
			case "+":
				power++
			case "-":
				if power > 0 {
					power--
				}
			default:
				{
					switch plan[index] {
					case "+":
						power++
					case "-":
						if power > 0 {
							power--
						}
					default:
					}
				}

			}
			result += power
			segment++
		}
	}
	return result
}

func part2(input string) {
	racetrack := `
S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-
`
	// 	input = `
	// A:+,-,=,=
	// B:+,=,-,+
	// C:=,-,+,+
	// D:=,=,=,+
	// `
	// 	racetrack = `
	// S+===
	// -   +
	// =+=-+
	// `
	parsed_racetrack := parse_racetrack(racetrack)
	plans := parse(input)
	results := make(map[string]int)
	for k, plan := range plans {
		results[k] = getResult(parsed_racetrack, plan, 10)
	}
	var ss []kv
	for k, v := range results {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s", kv.Key)
	}
	fmt.Printf("\n") // incorrect: BGEICDAHJ, ICDGJHAEB

}

func generatePlans() []string {
	var plans []string

	var f func(string, int, int, int)
	f = func(curr string, plus, minus, equals int) {
		if plus+minus+equals == 0 {
			plans = append(plans, curr)
		}

		if plus > 0 {
			f(curr+"+", plus-1, minus, equals)
		}

		if minus > 0 {
			f(curr+"-", plus, minus-1, equals)
		}

		if equals > 0 {
			f(curr+"=", plus, minus, equals-1)
		}
	}

	f("", 5, 3, 3)
	return plans
}

func parse_racetrack2(input string) []string {
	grid := parse_grid(input)
	n := len(grid)
	m := len(grid[0])
	var parsed []string
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	di := 0
	pos := []int{0, 1}
	for pos[0] != 0 || pos[1] != 0 {
		cur_c := grid[pos[0]][pos[1]]
		parsed = append(parsed, cur_c)
		for _, new_di := range []int{di, (di + 1) % 4, (di + 3) % 4} {
			cur_dir := directions[new_di]
			new_i := pos[0] + cur_dir[0]
			new_j := pos[1] + cur_dir[1]
			if 0 <= new_i && new_i < n && 0 <= new_j && new_j < m && grid[new_i][new_j] != " " {
				pos[0] = pos[0] + cur_dir[0]
				pos[1] = pos[1] + cur_dir[1]
				di = new_di
				break
			}
		}
	}
	parsed = append(parsed, grid[0][0])
	return parsed
}

func part3(input string) {
	racetrack := `
S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-
`

	plan := parse(input)["A"]
	// fmt.Println(plan)

	parsed_racetrack := parse_racetrack2(racetrack)
	// fmt.Println(parsed_racetrack)
	permutations := generatePlans()
	// fmt.Println(permutations)

	opponent_result := getResult(parsed_racetrack, plan, 2024)
	// fmt.Println("Result to beat", opponent_result)
	count := 0
	for i, pplan := range permutations {
		if (i+1)%500 == 0 {
			fmt.Println("Checking plan", i+1, "of", len(permutations))
		}
		plan = strings.Split(pplan, "")
		result := getResult(parsed_racetrack, plan, 2024)
		if result > opponent_result {
			count++
		}
	}
	fmt.Println(count)
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
