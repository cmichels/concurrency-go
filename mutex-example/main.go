package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
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



// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
//   m.Lock() 
// 	msg = s
//   m.Unlock() 
// }
// func main() {
//
//
//
// 	msg = "hello, worlds"
//   var mutex sync.Mutex
//
// 	wg.Add(2)
// 	go updateMessage("hello one", &mutex)
// 	go updateMessage("hello two", &mutex)
//
// 	wg.Wait()
//
// 	fmt.Println("msg:", msg)
//
// }
