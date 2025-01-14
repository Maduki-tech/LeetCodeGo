package main

import (
	"strings"
)

func reverseWords(s string) string {
	var result []string

	ssplitW := strings.Split(s, " ")

	for _, word := range ssplitW {

		left, right := 0, len(word)-1
		sWord := strings.Split(word, "")
		for left <= right {
			sWord[left], sWord[right] = sWord[right], sWord[left]

			left++
			right--
		}

		result = append(result, strings.Join(sWord, ""))
	}

	return strings.Join(result, " ")

}
