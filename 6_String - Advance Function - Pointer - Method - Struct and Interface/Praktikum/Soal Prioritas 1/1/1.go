package main

import (
	"fmt"
	"strings"
)

func main() {
  fmt.Println(pickVocals("alterra academy"))     // aea aae
  fmt.Println(pickVocals("sepulsa mantap jiwa")) // eua aa ia
  fmt.Println(pickVocals("kopi susu"))           // oi uu
}

func pickVocals(sentence string) string {
	vokal := "aeiou"
	result := ""
	kata := strings.Split(sentence, " ")

	for _, word := range kata {
		kataVokal := ""
		for _, char := range word {
			if strings.Contains(vokal, string(char)) {
				kataVokal += string(char)
			}
		}
		if kataVokal != "" {
			result += kataVokal + " "
		}
	}

	return strings.TrimSpace(result)
}