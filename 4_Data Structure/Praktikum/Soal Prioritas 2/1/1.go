package main

import "fmt"

func main() {
  fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    // [1,5,2,4,3]
  fmt.Println(spinSlice([]int{6, 7, 8}))          // [6,8,7]
  fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) // [8,1,5,2,4,3]
}

func spinSlice(numbers []int) []int {
  n := len(numbers)
  result := make([]int, 0)
  left, right := 0, n-1
  for left <= right {
    if left <= right {
      result = append(result, numbers[left])
      left++
    }
    if left <= right {
      result = append(result, numbers[right])
      right--
    }
  }
  return result
}