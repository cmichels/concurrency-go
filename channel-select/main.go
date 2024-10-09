package main

import (
	"fmt"
	"time"
)

func serverOne(ch chan string) {

	for {
		time.Sleep(6 * time.Second)
		ch <- "this is from sever one"
	}
}

func serverTwo(ch chan string) {

	for {
		time.Sleep(3 * time.Second)
		ch <- "this is from sever two"
	}
}

func main() {

	fmt.Println("select with channels")
	fmt.Println("--------------------")

	chOne := make(chan string)
	chTwo := make(chan string)

	go serverOne(chOne)
	go serverTwo(chTwo)

	for {
		select {
		case s1 := <-chOne:
			fmt.Println("case 1:", s1)
		case s2 := <-chOne:
			fmt.Println("case 2:", s2)
		case s3 := <-chTwo:
			fmt.Println("case 3: ", s3)
		case s4 := <-chTwo:
			fmt.Println("case 4: ", s4)
		default:
			// used to avoid deadlock
      fmt.Println("nothing listening")
		}
	}

}
