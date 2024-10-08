package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "plato", leftFork: 4, rightFork: 0},
	{name: "socrates", leftFork: 0, rightFork: 1},
	{name: "aristotle", leftFork: 1, rightFork: 2},
	{name: "pascal", leftFork: 2, rightFork: 3},
	{name: "locke", leftFork: 3, rightFork: 4},
}

var hunger = 3 // number of times they eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {

	fmt.Println("dining philosopers problem")
	fmt.Println("--------------------------")
	fmt.Println("table is empty")

	dine()

	fmt.Println("the table is empty")

}

func dine() {
	wg := &sync.WaitGroup{}

	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}

	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher,
	wg *sync.WaitGroup,
	forks map[int]*sync.Mutex,
	seated *sync.WaitGroup) {
	defer wg.Done()

}
