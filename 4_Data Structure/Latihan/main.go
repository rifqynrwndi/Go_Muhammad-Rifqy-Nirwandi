package main

import "fmt"

func main() {
	hasil, message := penjumlahan(2,3)
	fmt.Println(hasil, message)
}

func penjumlahan(a, b int)(int, string) {
	return a + b, "berhasil"

}