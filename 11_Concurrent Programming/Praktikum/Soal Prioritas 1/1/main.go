package main

import (
	"fmt"
	"time"
)

func reverseWord(word string, done chan bool) {
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
		fmt.Println(reversed)
		time.Sleep(3 * time.Second)
	}
	done <- true
}

func main() {
	word := "kasur"
	done := make(chan bool)

	go reverseWord(word, done)

	<-done
}
