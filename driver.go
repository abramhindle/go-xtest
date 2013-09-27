package main

import (
//	"fmt"
//	"os"
	"xtest"
)


func main() {
	disp := xtest.XOpenDisplay()
	defer xtest.XCloseDisplay(disp)
	for i := 0 ; i < 100 ; i += 1 {
		PressMouseButton(disp, 0)	
		ReleaseMouseButton(disp, 0)
	}
}
