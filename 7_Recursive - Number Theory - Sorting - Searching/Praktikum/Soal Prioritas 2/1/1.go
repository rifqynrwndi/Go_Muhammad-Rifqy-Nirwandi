package main

import (
  "fmt"
)

func main() {
  fmt.Println(catalan(7))  // 429
  fmt.Println(catalan(10)) // 16796
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func catalan(n int) int {
	pembilang := factorial(2*n)
	penyebut := factorial(n+1) * factorial(n)
	return pembilang / penyebut
}