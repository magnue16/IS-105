package main_test

func main() {
  str := "\x48\x65\x6C\x6C\x6F\x3A\x29"
  for i := 0; i < len(str); i++ {
    for j := 0; i < len (ASCII); j++ {
      if str[i] != ASCII[j] {
        fatal()
      }
    }
  }
}
