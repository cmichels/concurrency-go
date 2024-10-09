package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch

		fmt.Println("recieved ", i, " from channel")
		time.Sleep(1 * time.Second)
	}
}
func main() {

	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("sedning i ", i, " to chan")
		ch <- i
		fmt.Println("sent ", i, " to chan")

	}

	fmt.Println("done")
	close(ch)

}
