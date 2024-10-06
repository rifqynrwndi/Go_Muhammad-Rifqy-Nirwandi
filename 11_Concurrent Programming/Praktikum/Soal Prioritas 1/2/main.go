package main

import (
	"fmt"
	"math"
)

func generatePrimes(ch chan int) {
	for num := 2; num <= 100; num++ {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ch <- num
		}
	}
	close(ch)
}

func calculatePower(primeCh chan int, resultCh chan int) {
	for prime := range primeCh {
		resultCh <- prime * prime
	}
	close(resultCh)
}

func main() {
	primeCh := make(chan int)
	resultCh := make(chan int)

	go generatePrimes(primeCh)
	go calculatePower(primeCh, resultCh)

	for result := range resultCh {
		fmt.Println(result)
	}
}
