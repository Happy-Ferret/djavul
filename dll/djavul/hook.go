//+build djavul

package main

import "C"

import (
	"fmt"

	"github.com/sanctuary/djavul/d1/automap"
	"github.com/sanctuary/djavul/d1/capture"
	"github.com/sanctuary/djavul/d1/diablo"
	// TODO: update to github.com/AllenDang/w32 when
	// https://github.com/AllenDang/w32/pull/81 is merged.
	"github.com/TheTitanrain/w32"
)

//export Init
func Init() {
	// Called from WinMain.
	fmt.Println("init: djavul.dll")
}

// winPos specifies the window position and size.
//
//    0 = left  (  0, 0, 320, 480)
//    1 = right (320, 0, 320, 480)
//    2 = none  (  0, 0,   0,   0)
//    3 = max   (  0, 0, 640, 480)
var winPos int

//export OnKeyPress
func OnKeyPress(key int) {
	// Called from diablo_on_key_press.
	fmt.Println("key press:", key)
	switch key {
	case 'X':
		// eXplore
		automap.Explore()
	case 'Y':
		// Switch between different window positions and sizes.
		switch winPos {
		case 0:
			fmt.Println("window position: left")
			w32.MoveWindow(*diablo.Window, 0, 0, 320, 480, true)
			winPos++
		case 1:
			fmt.Println("window position: right")
			w32.MoveWindow(*diablo.Window, 320, 0, 320, 480, true)
			winPos++
		case 2:
			fmt.Println("window position: none")
			w32.MoveWindow(*diablo.Window, 0, 0, 0, 0, true)
			winPos++
		case 3:
			fmt.Println("window position: max")
			w32.MoveWindow(*diablo.Window, 0, 0, 640, 480, true)
			winPos = 0
		}
	}
}

//export OnKeyRelease
func OnKeyRelease(key int) {
	// Called from diablo_on_key_release.
	fmt.Println("key release:", key)
	switch key {
	case w32.VK_APPS, w32.VK_SNAPSHOT:
		// Capture screenshot.
		capture.Screenshot()
	}
}
