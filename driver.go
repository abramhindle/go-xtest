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
		xtest.PressMouseButton(disp, 1)	
		xtest.ReleaseMouseButton(disp, 1)
	}
}
