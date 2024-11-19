package main

import (
	"fmt"
	"slices"
	"strings"
)

func transpose(input []string) []string {
	matrix := make([][]string, len(input))
	for i, sentence := range input {
		matrix[i] = strings.Split(sentence, "")
	}
	xl := len(matrix[0])
	yl := len(matrix)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	var joined_result []string
	for _, sentence := range result {
		joined_result = append(joined_result, strings.Join(sentence, ""))
	}
	return joined_result
}

func find_occurences(sentences []string, set map[string]struct{}, words []string, transposed bool) map[string]struct{} {
	for lnr, sentence := range sentences {
		for _, word := range words {
			// fmt.Println("Searching for", word, "within", sentence)
			start := 0
			for len(sentence[start:]) > 0 {
				idx := strings.Index(sentence[start:], word)
				if idx == -1 {
					break
				}
				// fmt.Println("Found", word, "at", (start+idx)%(len(sentence)/2), "to", (start+idx+len(word)-1)%(len(sentence)/2))
				for i := range word {
					if transposed {
						set[fmt.Sprintf("%v,%v", start+idx+i, lnr)] = struct{}{}
					} else {
						// In the non transposed case we also duplicated each sentence
						set[fmt.Sprintf("%v,%v", lnr, (start+idx+i)%(len(sentence)/2))] = struct{}{}
					}
				}
				start += idx + 1
			}

		}
		// fmt.Println("Runes in sentence", len(set))
	}
	return set
}

func part3(input string) {
	// 	input = `WORDS:THE,OWE,MES,ROD,RODEO
	//
	// HELWORLT
	// ENIGWDXL
	// TRODEOAL`
	input_p := strings.Split(input, "\n\n")
	words := strings.Split(strings.Split(input_p[0], ":")[1], ",")
	for _, word := range words {
		if len(word) == 1 {
			continue
		}
		reversed := ""
		for _, v := range word {
			reversed = string(v) + reversed
		}
		if !slices.Contains(words, reversed) {
			words = append(words, reversed)
		}

	}
	sentences := strings.Split(strings.Trim(strings.Trim(strings.Trim(input_p[1], "\n"), " "), "\t"), "\n")
	for i, sentence := range sentences {
		sentences[i] += sentence
	}

	set := map[string]struct{}{}
	set = find_occurences(sentences, set, words, false)
	set = find_occurences(transpose(sentences), set, words, true)
	fmt.Println(len(set))
}
