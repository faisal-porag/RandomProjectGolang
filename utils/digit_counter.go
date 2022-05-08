package utils

import "fmt"

func DigitCountExample(){
   var n int
   fmt.Print("Enter the number: ")
   fmt.Scanf("%d", &n)
   count := 0
   for n >0 {
      n = n/10
      count++
   }
   fmt.Printf("The number of digits in the given number is: %d", count)
}
