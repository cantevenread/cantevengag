package main

import (
	"fmt"

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
		if timer == 0 {
			fmt.Println("cantevengag: error finding timer")
		}
		
	}

}
