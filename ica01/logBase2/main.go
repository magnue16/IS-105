package main

import (
  "./log"
  "fmt"
)

func main() {
  var (
    number float64 = 32
    log float64 = log.Log2(number)
  )
  fmt.Printf("The binary logarithm of %v is %f", number, log)
}
