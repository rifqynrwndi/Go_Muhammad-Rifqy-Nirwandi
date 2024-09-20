package main

import "fmt"

func main() {
  fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8})) // [[2,3,5,7],4,6,8]
  fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))     // [[3,5],15,24,9]
  fmt.Println(groupPrime([]int{4, 8, 9, 12}))         // [4,8,9,12]
}

func groupPrime(numbers []int) []any {
  primes := []int{}
	nonPrimes := []int{}

	for _, num := range numbers {
		if isPrime(num) {
			primes = append(primes, num)
		} else {
			nonPrimes = append(nonPrimes, num)
		}
	}

	if len(primes) == 0 {
		result := make([]any, len(nonPrimes))
		for i, v := range nonPrimes {
			result[i] = v
		}
		return result
	}

	result := make([]any, 0)
	result = append(result, primes)
	for _, num := range nonPrimes {
		result = append(result, num)
	}

	return result
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}