package main

import (
	"fmt"
	"math"
)

func generateComposites(ch chan int) {
	for num := 4; num <= 100; num++ {
		isComposite := false
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				isComposite = true
				break
			}
		}
		if isComposite {
			ch <- num
		}
	}
	close(ch)
}

func calculatePower(compositeCh chan int, resultCh chan int) {
	for composite := range compositeCh {
		resultCh <- composite * composite
	}
	close(resultCh)
}

func checkEvenOdd(resultCh chan int, doneCh chan bool) {
	for result := range resultCh {
		if result%2 == 0 {
			fmt.Println("Ping")
		} else {
			fmt.Println("Pong")
		}
	}
	doneCh <- true
}

func main() {
	compositeCh := make(chan int)
	resultCh := make(chan int)
	doneCh := make(chan bool)

	go generateComposites(compositeCh)
	go calculatePower(compositeCh, resultCh)
	go checkEvenOdd(resultCh, doneCh)

	<-doneCh
}
