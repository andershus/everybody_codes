package main

import (
	"fmt"
	"slices"
	"strings"
)

func part2(input string) {
	// 	input = `WORDS:THE,OWE,MES,ROD,HER,QAQ
	//
	// AWAKEN THE POWE ADORNED WITH THE FLAMES BRIGHT IRE
	// THE FLAME SHIELDED THE HEART OF THE KINGS
	// POWE PO WER P OWE R
	// THERE IS THE END
	// QAQAQ`
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
	// fmt.Println(words)
	sentences := strings.Split(input_p[1], "\n")
	// fmt.Println(sentences)

	set := map[string]struct{}{}
	for lnr, sentence := range sentences {
		for _, word := range words {
			// fmt.Println("Searching for", word, "within", sentence)
			start := 0
			for len(sentence[start:]) > 0 {
				idx := strings.Index(sentence[start:], word)
				if idx == -1 {
					break
				}
				// fmt.Println("Found", word, "at", start+idx, "to", start+idx+len(word)-1)
				for i := range word {
					set[fmt.Sprintf("%v,%v", lnr, start+idx+i)] = struct{}{}
				}
				start += idx + 1
				// fmt.Println("New subsentence", subsentence)
			}

		}
		// fmt.Println("Runes in sentence", len(set))
	}
	fmt.Println(len(set))
}

// // backwards
// reverse_sentence := ""
// for _, v := range sentence {
// 	reverse_sentence = string(v) + reverse_sentence
// }
// slen := len(sentence)
// fmt.Println("Going backwards through", reverse_sentence, "len:", slen, "looking for", word)
// subsentence = reverse_sentence
// start = 0
// for len(subsentence) > 0 {
// 	idx := strings.Index(subsentence, word)
// 	if idx == -1 {
// 		break
// 	}
// 	sidx := slen - start - idx
// 	fmt.Println("Found", word, "ending at", sidx)
// 	fmt.Println("which is", string(sentence[sidx-len(word):sidx]))
// 	for i := range word {
// 		fmt.Println("Adding", string(sentence[sidx-len(word)+i]))
// 		set[fmt.Sprintf("%v,%v", lnr, sidx-len(word)+i)] = struct{}{}
// 	}
// 	start = idx + 1
// 	subsentence = subsentence[start:]
// 	fmt.Println("New backwards subsentence", subsentence)
// 	// fmt.Println(subsentence)
// }
