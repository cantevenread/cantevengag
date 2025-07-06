package main

import (
	"fmt"
	"time"
	// "time"

	"github.com/cantevenread/cantevengag/gag"
)

func main() {
	gagChan := make(chan bool)
	go gag.GAGInit(gagChan)
	if <-gagChan {
		seedChan := make(chan bool)
		var timer int
		go func() { timer = gag.OpenSeedShop(seedChan) }()
		fmt.Println(<-seedChan)
		fmt.Println(timer)

		if timer > 0 {
			buySeeds := []string{"all"}
			buyFromMarket := make(chan bool)
			go gag.AnalyzeSeeds(buySeeds, 21, buyFromMarket)
			if <-buyFromMarket {
				homeChan := make(chan bool)
				go gag.GAGHome(homeChan)
				if <-homeChan {
					time.Sleep(time.Duration(timer) * time.Minute)

					seedChan2 := make(chan bool)
					var timer2 int
					go func() { timer2 = gag.OpenSeedShop(seedChan2) }()
					fmt.Println(<-seedChan2)
					fmt.Println(timer2)
					time.Sleep(time.Duration(timer2) * time.Minute)
				}
			}
		}
	}
}
