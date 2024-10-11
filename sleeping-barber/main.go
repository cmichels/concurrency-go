package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Yellow("the sleeping barber problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientChan:      clientChan,
		DoneChan:        doneChan,
		Open:            true,
	}

	color.Green("shop is open")

	shop.AddBarber("frank")
	shop.AddBarber("john")
	shop.AddBarber("tom")
	shop.AddBarber("timdog")
	shop.AddBarber("puppy")

	time.Sleep(5 * time.Second)

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShop()
		closed <- true

	}()

	i := 1
	go func() {
		for {
			randomMilli := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilli)):
				shop.addClient(fmt.Sprintf("client %d", i))
        i++

			}
		}
	}()

	<-closed
}
