package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j >= 6-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
