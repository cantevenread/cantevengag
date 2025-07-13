package gag

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cantevenread/cantevengag/img"
	"github.com/cantevenread/cantevengag/internal"
)

// takes player to home no matter where they are
// this closes all menus and resets them
func GAGHome(completed chan bool) {
	closeButtonChan := img.FindTemplateOnScreenAsync("./img/png/gag_x_button.png", 0.8)
	closeButton := <-closeButtonChan
	if closeButton.Err != nil {
		fmt.Println("cantevengag: close button not found", closeButton.Err)
	} else {
		fmt.Println("cantevengag: close button found")
		moved := make(chan bool)
		go internal.MoveMouseTo(closeButton.Coord.X, closeButton.Coord.Y, moved)
		if <-moved {
			internal.ClickMouse(nil)
		}
	}

	gardenButtonChan := img.FindTemplateOnScreenAsync("./img/png/gag_garden.png", 0.8)
	gardenButton := <-gardenButtonChan
	if gardenButton.Err != nil {
		fmt.Println("cantevengag: garden button not found", gardenButton.Err)
		completed <- false
	} else {
		fmt.Println("cantevengag: garden button found")
		moved := make(chan bool)
		go internal.MoveMouseTo(gardenButton.Coord.X, gardenButton.Coord.Y, moved)
		if <-moved {
			internal.ClickMouse(nil)
			fmt.Println("cantevengag: returned to garden")
			completed <- true
		}
	
	}
}

func GAGInit(completion chan bool) {
	windowCheck := make(chan bool)
	go internal.CheckIfWindowExists("RobloxPlayer", windowCheck)

	if !<-windowCheck {
		fmt.Println("cantevengag: roblox not found")
		completion <- false
		return
	}

	windowActivated := make(chan bool)
	go internal.ActivateWindow("RobloxPlayer", windowActivated)
	<-windowActivated
	internal.ClickMouse(nil)

	if internal.CurrentWindowFoucused() != "Roblox" {
		fmt.Println("Panic: Roblox did not focus as expected")
		completion <- false
		return
	}
	fmt.Println("cantevengag: roblox focused")
	// bug fix: fixes roblox focusing on the wrong window
	go internal.ActivateWindow("RobloxPlayer", nil)

	time.Sleep(200 * time.Millisecond)

	verifedgagChan := img.FindTemplateOnScreenAsync("./img/png/gag_garden.png", 0.8)
	verifedgag := <-verifedgagChan

	if verifedgag.Err != nil {
		fmt.Println("cantevengag: gag not found", verifedgag.Err)
		completion <- false
		return
	}
	fmt.Println("cantevengag: gag found")
	moved := make(chan bool)
	go internal.MoveMouseTo(verifedgag.Coord.X, verifedgag.Coord.Y, moved)
	if <-moved {
		internal.ClickMouse(nil)
	}

	// deactivate chat and leaderboard
	chatChan := img.FindTemplateOnScreenAsync("./img/png/rblx_chat_on.png", 0.8)
	chat := <-chatChan
	if chat.Completed {
		fmt.Println("cantevengag: chat is on, turning it off")
		internal.MoveMouseTo(chat.Coord.X, chat.Coord.Y, nil)
		time.Sleep(100 * time.Millisecond)
		internal.ClickMouse(nil)
	} else {
		fmt.Println("cantevengag: chat is off")
	}

	time.Sleep(100 * time.Millisecond)
	leaderboardChan := img.FindTemplateOnScreenAsync("./img/png/rblx_leaderboard_on.png", 0.8)
	leaderboard := <-leaderboardChan
	if leaderboard.Completed {
		fmt.Println("cantevengag: leaderboard is on, turning it off")
		internal.PressKey("tab", nil) // this toggles the leaderboard
		time.Sleep(100 * time.Millisecond)
	} else {
		fmt.Println("cantevengag: leaderboard is off")
	}






	// check if hotbar number 2 slot is emopty

	internal.PressKey("`", nil)

	// check is search is open

	secondSlotChan := img.FindTemplateOnScreenAsync("./img/png/gag_empty_slot2.png", 0.7)
	secondSlot := <-secondSlotChan
	if secondSlot.Completed {
		fmt.Println("cantevengag: hotbar slot 2 is empty")
	} else {
		fmt.Println("cantevengag: hotbar slot 2 is not empty, clearing it")
		go internal.ClickDragTo(620, 1060, 525, 820, nil)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("cantevengag: hotbar slot 2 cleared")

	}

	searchChan := img.FindTemplateOnScreenAsync("./img/png/gag_search.png", 0.8)
	search := <-searchChan
	if search.Err != nil {
		fmt.Println("cantevengag: search button not found", search.Err)
		completion <- false
	}
	fmt.Println("cantevengag: search button found")
	typeSearch := make(chan bool)
	go internal.MoveMouseAndClick(search.Coord.X, search.Coord.Y, typeSearch)
	if !<-typeSearch {
		completion <- false
	}
	fmt.Println("cantevengag: search button clicked ... searching recall wrench")
	time.Sleep(100 * time.Millisecond)
	recallWrenchChan := make(chan bool)
	go internal.KeyboardType("recall wrench", recallWrenchChan)
	if !<-recallWrenchChan {
		completion <- false
	}
	fmt.Println("cantevengag: recall wrench searched")
	isRecallInInvChan := img.FindTemplateOnScreenAsync("./img/png/gag_recall_wrench.png", 0.8)
	isRecallInInv := <-isRecallInInvChan
	if isRecallInInv.Err != nil {
		fmt.Println("cantevengag: recall wrench not found in inventory", isRecallInInv.Err)
		completion <- false
		return
	}
	fmt.Println("cantevengag: recall wrench found in inventory")
	internal.MoveMouseTo(isRecallInInv.Coord.X, isRecallInInv.Coord.Y, nil)
	time.Sleep(100 * time.Millisecond)
	internal.DoubleClickMouse(nil)
	fmt.Println("cantevengag: recall wrench placed in hotbar 2")
	// resetting search bar

	resetSearchChan := img.FindTemplateOnScreenAsync("./img/png/gag_x_inv.png", 0.8)
	resetSearch := <-resetSearchChan
	if resetSearch.Err != nil {
		fmt.Println("cantevengag: reset search button not found", resetSearch.Err)
		completion <- false
		return
	} else {
		fmt.Println("cantevengag: reset search button found")
		moved := make(chan bool)
		go internal.MoveMouseTo(resetSearch.Coord.X, resetSearch.Coord.Y, moved)
		if <-moved {
			internal.DoubleClickMouse(nil)
		}
	}
	time.Sleep(100 * time.Millisecond)
	internal.PressKey("escape", nil) // close inventory
	time.Sleep(100 * time.Millisecond)
	internal.PressKey("`", nil) // close search
	internal.PressKey("2", nil) // close hotbar
	time.Sleep(100 * time.Millisecond)
	completion <- true
}

func AnalyzeSeeds(seedsToPurchase []string, numberOfSeed int, completion chan bool) {
	// number of seed excludes carrot

	buySeed := func(bought chan bool) {
		internal.PressKey("down", nil)
		internal.PressKey("enter", nil)
		internal.PressKey("down", nil)
		time.Sleep(100 * time.Millisecond)
		for i := 0; i < 30; i++ {
			time.Sleep(50 * time.Millisecond)
			internal.PressKey("enter", nil)
		}
		internal.PressKey("up", nil)
		internal.PressKey("enter", nil)
		bought <- true
	}

	// for loop over each seed in seedtoPurchase Array
	for _, seed := range seedsToPurchase {
		switch seed {
		case "all":
			fmt.Println("cantevengag: buying all seeds")
			navagationMode := make(chan bool)
			go internal.PressKey("\\", navagationMode)
			if <-navagationMode {
				// initial navigation
				internal.PressKey("down", nil)
				internal.PressKey("down", nil)
				internal.PressKey("enter", nil)
				internal.PressKey("down", nil)
				for i := 0; i < 30; i++ {
					time.Sleep(25 * time.Millisecond)
					internal.PressKey("enter", nil)
				}
				internal.PressKey("up", nil)
				internal.PressKey("enter", nil)

				time.Sleep(400 * time.Millisecond)
				for i := 0; i < numberOfSeed; i++ {
					bought := make(chan bool)
					go buySeed(bought)
					<-bought
					time.Sleep(70 * time.Millisecond)
				}
				fmt.Println("cantevengag: going back to top")

				for i := 0; i < numberOfSeed; i++ {
					internal.PressKey("up", nil)
					time.Sleep(100 * time.Millisecond)
				}

				internal.PressKey("\\", nil) // exit navigation mode
				fmt.Println("cantevengag: all seeds purchased, exiting navigation mode")
				completion <- true
				return

				// gagHome
			}
		}
	}
	completion <- false
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

	seedSellerChan := img.FindTemplateOnScreenAsync("./img/png/gag_seed_shop.png", 0.65)

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
	verifySeedShopChan := img.FindTemplateOnScreenAsync("./img/png/gag_verify_seed_shop.png", 0.62)
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
		channels[label] = img.FindTemplateOnScreenAsync(path, 0.8)
	}

	var timerValue int
    var timerFound bool
	
	// Normal check
    for label, ch := range channels {
        result := <-ch
        // fmt.Printf("%s: %v\n", label, result)
        if result.Completed {
            timerFound = true
            if labelInt, err := strconv.Atoi(label); err == nil {
                timerValue = 1 + labelInt
                fmt.Printf("cantevengag: seed shop timer is %s minutes\n", label)
            } else {
                fmt.Printf("Error converting label to int: %v\n", err)
            }
            break
        } else {
            fmt.Println("cantevengag: seed shop timer is not " + label + " minutes, trying next timer")
        }
    }

    if !timerFound {
		time.Sleep(3 * time.Second)
        // Re-check
        for label, path := range timerPaths {
            recheckChan := img.FindTemplateOnScreenAsync(path, 0.8)
            recheckResult := <-recheckChan
            if recheckResult.Completed {
                timerFound = true
                if labelInt, err := strconv.Atoi(label); err == nil {
                    timerValue = 1 + labelInt
                    fmt.Printf("cantevengag: re-checked seed shop timer is %s minutes\n", label)
                } else {
                    fmt.Printf("Error converting label to int: %v\n", err)
                }
                break
            } 
        }
    }

	if timerFound {
        completion <- true
        return timerValue
    } else {
        fmt.Println("cantevengag: seed shop timer <1 minute, timer set to 2 minutes")
        completion <- true
        return 2
    }

}
