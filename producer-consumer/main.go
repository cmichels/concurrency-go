package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	numberOfPizzas = 10
)

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	orderNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch

	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	for {
     
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("the pizzeria is open for business")
	color.Cyan("---------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

}
