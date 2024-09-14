package main

import "fmt"

func main() {
	var budget int
	var waktu int
	var sulit int
	var skor int

	fmt.Printf("Masukkan budget: ")
	fmt.Scanln(&budget)
	
	fmt.Printf("Masukkan waktu pengerjaan: ")
	fmt.Scanln(&waktu)
	
	fmt.Printf("Masukkan tingkat kesulitan: ")
	fmt.Scanln(&sulit)

	switch {
		case budget >= 0 && budget <= 50:
			skor += 4
		case budget >= 51 && budget <= 80:
			skor += 3
		case budget >= 81 && budget <= 100:
			skor += 2
		case budget > 100:
			skor += 1
	}

	switch {
	case waktu >= 0 && waktu <= 20:
		skor += 10
	case waktu >= 21 && waktu <= 30:
		skor += 5
	case waktu >= 31 && waktu <= 50:
		skor += 2
	case waktu > 50:
		skor += 1
	}

	switch {
	case sulit >= 0 && sulit <= 3:
		skor += 10
	case sulit >= 4 && sulit <= 6:
		skor += 5
	case sulit >= 8 && sulit <= 10:
		skor += 1
	case sulit > 10:
		skor += 0
	}

	fmt.Print("Output: ")

	switch {
	case skor >= 17 && skor <= 24:
		fmt.Print("High")
	case skor >= 10 && skor <= 16:
		fmt.Print("Medium")
	case skor >= 3 && skor <= 9:
		fmt.Print("Low")
	case skor <= 2:
		fmt.Print("Impossible")
	}
}
