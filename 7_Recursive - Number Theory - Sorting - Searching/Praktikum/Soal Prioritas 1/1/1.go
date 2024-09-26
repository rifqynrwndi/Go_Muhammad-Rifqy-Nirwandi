package main

import "fmt"

func main() {
  fmt.Println(fibX(5))  // 12
  fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	if n < 0 {
		return 0
	}

	a,b := 0, 1
	sum := a

	for i := 1; i <= n; i++ {
		sum += b
		a, b = b, a+b
	}

	return sum
}