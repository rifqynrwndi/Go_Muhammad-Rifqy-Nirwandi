package main

import "fmt"

func main() {
   res1 := merge([][]int{
       {1, 2, 5, 4, 3},
       {1, 2, 7, 8},
   })
   fmt.Println(res1) 

   res2 := merge([][]int{
       {1, 2, 5, 4, 5},
       {7, 9, 10, 5},
   })
   fmt.Println(res2) 

   res3 := merge([][]int{
       {1, 4, 5},
       {7},
   })
   fmt.Println(res3) 
}

func merge(data [][]int) []int {
    unique := make(map[int]bool)
  
    for _, arr := range data {
      for _, val := range arr {
        unique[val] = true
      }
    }
    
    result := []int{}
    for key := range unique {
      result = append(result, key)
    }
  
    return result
}