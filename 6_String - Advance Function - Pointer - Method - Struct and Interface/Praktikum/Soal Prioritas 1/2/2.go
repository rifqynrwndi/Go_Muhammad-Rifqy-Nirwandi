package main

import "fmt"

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	result := ""
	left := 0 
	right := len(sentence)-1

	for left <= right {
		if left == right {
			result += string(sentence[left])
		} else {
			result += string(sentence[left]) + string(sentence[right])
		}
		left++
		right--
	}

	return result
}
