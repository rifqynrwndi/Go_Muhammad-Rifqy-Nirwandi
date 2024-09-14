package main

import "fmt"

func hitungJumlah(kata string) map[rune]int {
	jumlah := make(map[rune]int)
	for _, huruf := range kata {
		if huruf != ' ' {
			jumlah[huruf]++
		}
	}
	return jumlah
}

func isAnagram(kata1, kata2 string) bool {
	if len(kata1) != len(kata2) {
		return false
	}

	hitungJumlah1 := hitungJumlah(kata1)
	hitungJumlah2 := hitungJumlah(kata2)

	for huruf, jumlah := range hitungJumlah1 {
		if hitungJumlah2[huruf] != jumlah {
			return false
		}
	}
	return true
}

func main() {
	var kata1, kata2 string

	fmt.Print("Kata pertama: ")
	fmt.Scanln(&kata1)

	fmt.Print("Kata kedua: ")
	fmt.Scanln(&kata2)

	if isAnagram(kata1, kata2) {
		fmt.Println("Anagram")
	} else {
		fmt.Println("Bukan Anagram")
	}
}
