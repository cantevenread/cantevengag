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
	robotgo.TypeStr(text)
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