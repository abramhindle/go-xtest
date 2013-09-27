package main

import (
	"fmt"
//	"os"
	"xtest"
)


func main() {
	disp := xtest.XOpenDisplay()
	screen := xtest.XGetScreen(disp)
	defer xtest.XCloseDisplay(disp)
	for i := 0 ; i < 100 ; i += 1 {
		xtest.PressMouseButton(disp, 1)	
		xtest.ReleaseMouseButton(disp, 1)
		xtest.MoveMouseAbs(disp, screen, 100+i,100+i)
		xtest.Usleep(1000);
		color := xtest.GetPixelColor(disp, 100+i,100+i)
		fmt.Printf("%v\n", color)
	}
}
