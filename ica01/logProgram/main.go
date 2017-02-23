package main

import (
  "./logbcli"
  "fmt"
  "os"
  "strconv"
)

func main() {

    arg1 := os.Args[1]
    arg2 := os.Args[2]
    argInput1, err := strconv.ParseFloat(arg1, 64)
    argInput2, err := strconv.ParseFloat(arg2, 64)
    if argInput1 == 2 {
      log1 := log.Log2(argInput2)
      fmt.Printf("The binary logarithm of %v is %f", arg2, log1)
    }
    if argInput1 == 10 {
      log2 := log.Log10(argInput2)
      fmt.Printf("The logarithm of %v in base 10 is %f", arg2, log2)
    }
    if err != nil {
      fmt.Println(err)
    }
}
