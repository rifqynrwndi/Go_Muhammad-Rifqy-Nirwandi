package main

import "fmt"

func main() {
   fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
   fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
   fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
   sum := 0
   for _, num := range numbers {
     if isPrime(num) {
       sum += num
     }
   }
   return sum
 }
 
 func isPrime(n int) bool {
   if n < 2 {
     return false
   }
   for i := 2; i*i <= n; i++ {
     if n%i == 0 {
       return false
     }
   }
   return true
 }