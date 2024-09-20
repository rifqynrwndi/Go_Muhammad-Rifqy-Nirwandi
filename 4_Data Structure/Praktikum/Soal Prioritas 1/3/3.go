package main

import (
	"fmt"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("sum: %.2f\n", sum(data1))       // 71.00
	fmt.Printf("mean: %.2f\n", mean(data1))     // 5.46
	fmt.Printf("median: %.2f\n", median(data1)) // 5.00
	fmt.Printf("mode: %.2f\n", mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("sum: %.2f\n", sum(data2))       // 103.00
	fmt.Printf("mean: %.2f\n", mean(data2))     // 7.92
	fmt.Printf("median: %.2f\n", median(data2)) // 8.00
	fmt.Printf("mode: %.2f\n", mode(data2))     // 1.00
}

func sum(data []float64) float64 {
	total := 0.0
	for _, v := range data {
		total += v
	}
	return total
}

func mean(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func median(data []float64) float64 {
	sortedData := mergeSort(data)
	n := len(sortedData)
	if n%2 == 0 {
		return (sortedData[n/2-1] + sortedData[n/2]) / 2
	}
	return sortedData[n/2]
}

func mode(data []float64) float64 {
	frequency := make(map[float64]int)
	maxFreq := 0
	mode := data[0]
	for _, v := range data {
		frequency[v]++
		if frequency[v] > maxFreq {
			maxFreq = frequency[v]
			mode = v
		}
	}
	return mode
}

func mergeSort(data []float64) []float64 {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := mergeSort(data[:mid])
	right := mergeSort(data[mid:])

	return merge(left, right)
}

func merge(left, right []float64) []float64 {
	result := []float64{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
