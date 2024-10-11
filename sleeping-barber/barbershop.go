package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	DoneChan        chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) AddBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s checks for clients", barber)

		for {
			if len(shop.ClientChan) == 0 {
				color.Yellow("no clients for %s. nap time", barber)
				isSleeping = true
			}

			client, open := <-shop.ClientChan

			if open {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
          isSleeping = false
				}

				shop.cutHair(barber, client)
			} else {
				shop.sendBarberHope(barber)
				return
			}
		}
	}()

}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s hair", barber, client)
}
func (shop *BarberShop) sendBarberHope(barbert string) {
	color.Cyan("%s is going home", barbert)
	shop.DoneChan <- true
}
func (shop *BarberShop) CloseShop() {
	color.Cyan("closing shop")
	close(shop.ClientChan)
	shop.Open = false

	for i := 0; i < shop.NumberOfBarbers; i++ {
		<-shop.DoneChan
	}

	close(shop.DoneChan)

	color.Green("barber shop is closed")
	color.Green("---------------------")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("***** %s arrives", client)
	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Yellow("%s sits in waiting room", client)
		default:
			color.Red("%s leaves", client)
		}

	} else {
		color.Red("shop is closed %s leaves", client)
	}
}
