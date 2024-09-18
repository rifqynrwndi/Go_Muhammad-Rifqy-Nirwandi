package main

import "fmt"

func main() {
	var nilai int

	fmt.Printf("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else if nilai >= 85 {
		fmt.Println("Nilai A")
	} else if nilai >= 70 {
		fmt.Println("Nilai B")
	} else if nilai >= 55 {
		fmt.Println("Nilai C")
	} else if nilai >= 40 {		
		fmt.Println("Nilai D")
	} else if nilai >= 0 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}