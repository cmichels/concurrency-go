package main

import (
	"fmt"
	"time"
)


func printSomthing(s string)  {
  fmt.Println(s)
}
func main() {
  
  go printSomthing("first thing to print")

  time.Sleep(1 * time.Second)
  printSomthing("second to print")
}
