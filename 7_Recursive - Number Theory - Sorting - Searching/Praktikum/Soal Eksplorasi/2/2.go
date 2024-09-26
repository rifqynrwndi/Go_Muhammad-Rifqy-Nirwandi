package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(items)))
	var result []int
	for _, item := range items {
		if target-item >= 0 {
			result = append(result, item)
			target -= item
		}
		if target == 0 {
			break
		}
	}
	return result
}
