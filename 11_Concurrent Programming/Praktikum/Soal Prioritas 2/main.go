package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

func readFile(filename string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		close(ch)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			ch <- strings.ToLower(word)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file:", err)
	}

	close(ch)
}

func countWords(ch <-chan string, freqMap map[string]int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for word := range ch {
		mutex.Lock()
		freqMap[word]++
		mutex.Unlock()
	}
}

func saveToCSV(filename string, freqMap map[string]int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error membuat file CSV:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"word", "frequencies"})
	for word, freq := range freqMap {
		writer.Write([]string{word, fmt.Sprintf("%d", freq)})
	}
}

func main() {
	filename := "document.txt"
	ch := make(chan string)
	freqMap := make(map[string]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(1)
	go readFile(filename, ch, &wg)

	wg.Add(1)
	go countWords(ch, freqMap, &wg, &mutex)

	wg.Wait()

	saveToCSV("word_frequencies.csv", freqMap)

	fmt.Println("Proses selesai. Hasil disimpan di word_frequencies.csv")
}
