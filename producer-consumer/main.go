package main

import (
	"fmt"
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

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("received order %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("making pizza %d. ready in %d seconds\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** ran out of ingredients for pizza %d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** cook quit for %d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("pizza order %d is read", pizzaNumber)
		}

		p := PizzaOrder{
			orderNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaOrder{
		orderNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {

	var i = 0
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.orderNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
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

	for i := range pizzaJob.data {
		if i.orderNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("order %d out for delivery", i.orderNumber)
			} else {
				color.Red(i.message)
				color.Red("upset customer")
			}
		} else {
			color.Cyan("done making pizzas.....")
			err := pizzaJob.Close()

			if err != nil {
				color.Red("error closing change: %s", err)
			}
		}
	}

	color.Cyan("----------------")
	color.Cyan("done for the day")

	color.Cyan("made %d, failed %d, total %d", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("it was an awefule day")
	case pizzasFailed > 6:
		color.Red("not a good day")
	case pizzasFailed >= 4:
		color.Red("it was ok")
	case pizzasFailed >= 2:
		color.Yellow("pretty good day")
	default:
		color.Green("great day")

	}

}
