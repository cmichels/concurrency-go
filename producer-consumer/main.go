package main

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
  message string
  success bool
}

func main() {

}
