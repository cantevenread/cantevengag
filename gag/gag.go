package gag

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cantevenread/cantevengag/img"
	"github.com/cantevenread/cantevengag/internal"
)

func GAGInit(completion chan bool) {
	windowCheck := make(chan bool)
	go internal.CheckIfWindowExists("RobloxPlayer", windowCheck)

	if <-windowCheck {
		windowActivated := make(chan bool)
		go internal.ActivateWindow("RobloxPlayer", windowActivated)
		<-windowActivated
		go internal.ClickMouse(nil)

		if internal.CurrentWindowFoucused() != "Roblox" {
			fmt.Println("Panic: Roblox did not focus as expected")
			completion <- false
			return
		} else {
			fmt.Println("cantevengag: roblox focused")
			// bug fix: fixes roblox focusing on the wrong window
			go internal.ActivateWindow("RobloxPlayer", nil)
		}

	} else {
		fmt.Println("cantevengag: roblox not found")
		completion <- false
		return
	}

	// make sure gag is open
	time.Sleep(200 * time.Millisecond)

	verifedgagChan := img.FindTemplateOnScreenAsync("./img/png/gag_garden.png", 0.8)

	verifedgag := <-verifedgagChan

	if verifedgag.Err != nil {
		fmt.Println("cantevengag: gag not found", verifedgag.Err)
		completion <- false
		return
	} else {
		fmt.Println("cantevengag: gag found")
		moved := make(chan bool)
		go internal.MoveMouseTo(verifedgag.Coord.X, verifedgag.Coord.Y, moved)
		if <-moved {
			internal.ClickMouse(nil)
		}
		completion <- true

	}
}

func OpenSeedShop(completion chan bool) (timer int) {
	seedShopChan := img.FindTemplateOnScreenAsync("./img/png/gag_seeds.png", 0.8)
	seedShop := <-seedShopChan

	if seedShop.Err != nil {
		fmt.Println("cantevengag: seed shop teleport not found", seedShop.Err)
		completion <- false
	} else {
		fmt.Println("cantevengag: seed shop teleport found")
		moved := make(chan bool)
		go internal.MoveMouseTo(seedShop.Coord.X, seedShop.Coord.Y, moved)
		if <-moved {
			internal.ClickMouse(nil)
		}
	}

	seedSellerChan := img.FindTemplateOnScreenAsync("./img/png/gag_seed_shop.png", 0.4)

	seedSeller := <-seedSellerChan

	if seedSeller.Err != nil {
		fmt.Println("cantevengag: seed seller not found trying anyways:", seedSeller.Err)
		internal.PressKey("e", nil)
	} else {
		fmt.Println("cantevengag: seed seller found ")
		internal.MoveMouseTo(seedSeller.Coord.X, seedSeller.Coord.Y, nil)
		time.Sleep(1 * time.Second)
		internal.PressKey("e", nil)
	}

	time.Sleep(2 * time.Second)

	// make sure the seed shop is open
	verifySeedShopChan := img.FindTemplateOnScreenAsync("./img/png/gag_verify_seed_shop.png", 0.4)
	verifySeedShop := <-verifySeedShopChan
	if verifySeedShop.Err != nil {
		fmt.Println("cantevengag: seed shop not found", verifySeedShop.Err)
		completion <- false
		return 0
	} else {
		fmt.Println("cantevengag: seed shop found")
		internal.MoveMouseTo(verifySeedShop.Coord.X, verifySeedShop.Coord.Y, nil)
	}
	// timer
	timerPaths := map[string]string{
		"1": "./img/png/timer/gag_timer_1m.png",
		"2": "./img/png/timer/gag_timer_2m.png",
		"3": "./img/png/timer/gag_timer_3m.png",
		"4": "./img/png/timer/gag_timer_4m.png",
	}

	channels := make(map[string]<-chan img.FindResult)
	for label, path := range timerPaths {
		channels[label] = img.FindTemplateOnScreenAsync(path, 0.9)
	}

	allFalse := true
	for label, ch := range channels {
		result := <-ch
		// fmt.Printf("%s: %v\n", label, result)
		if result.Completed {
			allFalse = false
			if labelInt, err := strconv.Atoi(label); err == nil {
				fmt.Printf("cantevengag: seed shop timer is %s minutes\n", label)
				completion <- true
				return 1 + labelInt
			} else {
				fmt.Printf("Error converting label to int: %v\n", err)
				completion <- false
				return 0
			}
		} else {
			fmt.Println("cantevengag: seed shop timer is not " + label + " minutes, trying next timer")
		}
	}

	if allFalse {
		completion <- true
		fmt.Println("cantevengag: seed shop timer <1 minute, timer set to 1 minute")
		return 1
	}

	return 0
}
