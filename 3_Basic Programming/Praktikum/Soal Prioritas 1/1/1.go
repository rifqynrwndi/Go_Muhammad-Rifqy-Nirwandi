package main

import "fmt"

func main() {
	var v float32 = 0
	const phi float32 = 3.14
	var r float32
	var t float32

	fmt.Printf("Masukkan jari-jari: ")
	fmt.Scanln(&r) 
	
	fmt.Printf("Masukkan tinggi: ")
	fmt.Scanln(&t) 

	v = phi * r * r * t
	fmt.Printf("Volume tabung: %.2f\n", v)
}
