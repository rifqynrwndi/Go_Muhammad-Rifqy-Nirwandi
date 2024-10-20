package models

import (
	"strings"
)

type Word struct {
	Word        string `json:"word"`
	Length      int    `json:"length"`
	NumOfVocals int    `json:"num_of_vocals"`
	Palindrome  bool   `json:"palindrome"`
}

var Words = []Word{}

func CountVocals(word string) int {
	vocals := "aeiouAEIOU"
	count := 0
	for _, char := range word {
		if strings.ContainsRune(vocals, char) {
			count++
		}
	}
	return count
}

func IsPalindrome(word string) bool {
	processed := strings.ToLower(strings.ReplaceAll(word, " ", ""))
	reversed := reverse(processed)
	return processed == reversed
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
