package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("init balance: %d", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "main", Amount: 500},
		{Source: "gifts", Amount: 10},
		{Source: "part time", Amount: 50},
		{Source: "invest", Amount: 100},
	}

	wg.Add(len(incomes))
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("week %d, earned $%d from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("final balance: $%d", bankBalance)

}
