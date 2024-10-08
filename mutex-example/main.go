package main

import (
	"fmt"
	"sync"
)


var msg string
var wg sync.WaitGroup

func updateMessage(s string)  {
  defer wg.Done()
  msg = s
}
func main() {
  

  msg = "hello, worlds"

  wg.Add(2)
  go updateMessage("hello one")
  go updateMessage("hello two")

  wg.Wait()

  fmt.Println("msg:", msg)
  
}
