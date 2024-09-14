package main

import "fmt"

func main() {
	var nilai int

	fmt.Printf("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 85 && nilai <= 100 {
		fmt.Println("Nilai A")
	} else if nilai >= 70 && nilai <= 84 {
		fmt.Println("Nilai B")
	} else if nilai >= 55 && nilai <= 69 {
		fmt.Println("Nilai C")
	} else if nilai >= 40 && nilai <= 50 {		
		fmt.Println("Nilai D")
	} else if nilai >= 0 && nilai <= 39 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
