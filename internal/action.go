package internal

import (
	"time"

	"github.com/go-vgo/robotgo"
)

type Coordinate struct {
	X int
	Y int
}

func CurrentMousePosition() Coordinate {
	x, y := robotgo.Location()
	return Coordinate{x, y}
}

func MoveMouseTo(x int, y int, completed chan bool) {
	robotgo.Move(x, y)
	if completed != nil {
		completed <- true
	}
}

func ClickMouse(completed chan bool) {
	robotgo.Click()
	if completed != nil {
		completed <- true
	}
}

func MoveMouseAndClick(x int, y int, completed chan bool) {
	robotgo.Move(x, y)
	time.Sleep(25 * time.Millisecond) 
	robotgo.Click()
	if completed != nil {
		completed <- true
	}
}

func DoubleClickMouse(completed chan bool) {
	robotgo.Click()
	time.Sleep(300 * time.Millisecond)
	robotgo.Click()
	if completed != nil {
		completed <- true
	}
}

func GetAllWindows() ([]string, error) {
	return robotgo.FindNames()
}

func ActivateWindow(window string, completed chan bool) {
	robotgo.ActiveName(window)
	if completed != nil {
		completed <- true
	}
}

func KeyboardType(text string, completed chan bool) {
	// slice the string into  an array of single characters
	for _, char := range text {
		robotgo.KeyPress(string(char))
		time.Sleep(50 * time.Millisecond) // add a small delay between each character
	}

	if completed != nil {
		completed <- true
	}
}

func CheckIfWindowExists(window string, completed chan bool) {
	windows, _ := GetAllWindows()
	for _, w := range windows {
		if w == window {
			if completed != nil {
				completed <- true
			}
			return
		}
	}
	if completed != nil {
		completed <- false
	}
}

func CurrentWindowFoucused() string {
	return robotgo.GetTitle()
}

func PressKey(key string, completed chan bool) {
	robotgo.KeyTap(key)
	if completed != nil {
		completed <- true
	}
}

func ClickDragTo(x1, y1, x2, y2 int, completed chan bool) {
    // Move to starting point
    robotgo.Move(x1, y1)
    robotgo.MilliSleep(10)

    robotgo.Toggle("left")          // equivalent to pressing down :contentReference[oaicite:2]{index=2}
    robotgo.MilliSleep(50)

    robotgo.DragSmooth(x2, y2)      // smoothly move mouse while holding :contentReference[oaicite:3]{index=3}

    robotgo.Toggle("left", "up")    // release :contentReference[oaicite:4]{index=4}

    // Signal completion
    if completed != nil {
        completed <- true
    }
}
