package main

import (
	"fmt"
	"time"

	"github.com/cantevenread/cantevengag/gag"
)


func main() {

	gagInitDone := make(chan bool)
	go gag.GAGInit(gagInitDone)
	if !<-gagInitDone {
		fmt.Println("cantevengag: could not initialize")
		return
	}
	fmt.Println("cantevengag: initialized successfully")	

	for {
		seedChan := make(chan bool)
		var timer int
		go func() { timer = gag.OpenSeedShop(seedChan) }()
		if !<-seedChan {
			fmt.Println("❌ Failed to open seed shop.")
			continue
		}
		fmt.Printf("cantevengag:  Got timer: %d minutes ⏱️ \n", timer)

		analyzeChan := make(chan bool)
		go gag.AnalyzeSeeds([]string{"all"}, 21, analyzeChan)
		if !<-analyzeChan {
			fmt.Println("cantevengag: failed to analyze seeds.")
			continue
		}

		homeChan := make(chan bool)
		go gag.GAGHome(homeChan)
		if !<-homeChan {
			fmt.Println("cantevengag: failed to return home.")
			continue
		}

		fmt.Printf("cantevengag: ⌛ Waiting %d minutes...\n", timer)
		time.Sleep(time.Duration(timer) * time.Minute)
	}
}
