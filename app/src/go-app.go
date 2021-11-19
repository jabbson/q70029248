package main

import (
  "fmt"
  "time"
)

func main() {
  for i := 0; i < 1000; i ++ {
    fmt.Printf("Log Entry #%d\n", i)
    time.Sleep(time.Second * 1)
  }
}
