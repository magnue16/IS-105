package main

import (
  "./logcli"
  "fmt"
  "os"
  "strconv"
)

func main() {

    arg1 := os.Args[1]
    argInput, err := strconv.ParseFloat(arg1, 64)
    log := log.Log2(argInput)
    if err != nil {
      fmt.Println(err)
    }

  fmt.Printf("The binary logarithm of %v is %f", arg1, log)
}
