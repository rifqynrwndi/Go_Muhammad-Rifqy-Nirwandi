package main

import "fmt"

func main() {
  fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
  fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
  fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) []any {
	palin := []string{}
	nonPalin := []string{}
	for _, word := range words {
		if isPalindrome(word) {
			palin = append(palin, word)
		} else {
			nonPalin = append(nonPalin, word)
		}
	}
	result := []any{}
	if len(palin) > 0 {
		result = append(result, palin)
	}
	for _, word := range nonPalin {
		result = append(result, word)
	}
	return result
}

func isPalindrome(word string) bool {
	length := len(word)
	for i := 0; i < length/2; i++ {
		if word[i] != word[length-i-1] {
			return false
		}
	}
	return true
}