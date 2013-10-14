/* X11::GUITest ($Id: GUITest.xs 218 2013-02-01 18:29:03Z pecastro $)
 *
 * Copyright (c) 2003-2011  Dennis K. Paulsen, All Rights Reserved.
 * Email: ctrondlp@cpan.org
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation; either version 2 of
 * the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses>.
 *
 */
/* http://cpansearch.perl.org/src/PECASTRO/X11-GUITest-0.27/GUITest.xs */

package xtest

/*
#cgo pkg-config: x11 xtst
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <X11/Xlib.h>
#include <X11/Xutil.h>
#include <X11/Xatom.h>
#include <X11/Xlocale.h>
#include <X11/Intrinsic.h>
#include <X11/StringDefs.h>
#include <X11/keysym.h>
#include <X11/extensions/XTest.h>

// http://rosettacode.org/wiki/Color_of_a_screen_pixel
void get_pixel_color (Display *d, int x, int y, unsigned int * red, unsigned int * green, unsigned int * blue) {
  XColor color;
  XImage *image;
  image = XGetImage (d, RootWindow (d, DefaultScreen (d)), x, y, 1, 1, AllPlanes, XYPixmap);
  color.pixel = XGetPixel (image, 0, 0);
  XFree (image);
  XQueryColor (d, DefaultColormap(d, DefaultScreen (d)), &color);
  *red  = color.red;
  *green = color.red;
  *blue  = color.red;
}
 
*/
import "C"

import (
	//"errors"
	//"unsafe"
)

type XDisplay *C.Display

func XOpenDisplay() XDisplay {
	disp := C.XOpenDisplay(nil)
	return disp;
}

func XCloseDisplay(disp XDisplay) {
	C.XSync(disp, 0)
	C.XCloseDisplay(disp);
}

func XGetScreen(disp XDisplay) C.int {
	theScreen := C.XDefaultScreen(disp)
	C.XSync(disp, C.int(1))
	return theScreen;
}


func PressMouseButton(disp XDisplay, button int) C.int {
	retval := C.XTestFakeButtonEvent(disp, C.uint(button), C.int(1), C.ulong(1))
	return retval
}
func ReleaseMouseButton(disp XDisplay, button int) C.int {
	retval := C.XTestFakeButtonEvent(disp, C.uint(button), C.int(0), C.ulong(1))
	return retval
}

func MoveMouseAbs(disp XDisplay, screen C.int, x, y int) bool {
 	if (screen >= 0 && screen < C.XScreenCount(disp)) {
 		/* I decided not to set our error handler, since the
		 window must exist. */
 		C.XWarpPointer(disp, C.None,
 			C.XRootWindow(disp, screen),
 			     0, 0, 0, 0,
 			C.int(x), C.int(y));
 		C.XSync(disp, C.int(0));
		return true;
 	} 
	return false;
}

func Usleep(usec int) {
	C.usleep(C.__useconds_t(usec))
}

func GetKeycodeFromKeysym(display XDisplay, keysym C.KeySym) C.KeyCode {
	//-#define XK_Alt_L                         0xffe9  /* Left alt */
	//-#define XK_Meta_L                        0xffe7  /* Left meta */
 	kc := C.XKeysymToKeycode(display, keysym)
 	if (kc == 0 && keysym == 0xffe9) {
 		return C.XKeysymToKeycode(display, 0xffe7)
	}
	return kc
}

/* Function: PressKeyImp
 * Description: Presses the key for the specified keysym.  Lower-level
 * 				implementation.
 * Note: Returns TRUE (non-zero) on success, FALSE (zero) on failure.
 */
func PressKey(disp XDisplay, sym C.KeySym) bool  {
	kc := GetKeycodeFromKeysym(disp, sym)
	if (kc == 0) {
		return false
	}
	retval := C.XTestFakeKeyEvent(disp, C.uint(kc), 1, 10)
	C.XFlush( disp )
	return (retval != 0)
}

func ReleaseKey(disp XDisplay, sym C.KeySym) bool  {
	kc := GetKeycodeFromKeysym(disp, sym)
	if (kc == 0) {
		return false
	}
	retval := C.XTestFakeKeyEvent(disp, C.uint(kc), 0, 10)
	C.XFlush( disp )
	return (retval != 0)
}


// static Display *TheXDisplay = NULL;
// TheXDisplay = XOpenDisplay(NULL);
// 
// /* Function: CloseXDisplay
//  * Description: Closes our connection to the X server's display
//  */
// static void CloseXDisplay(void)
// {
// 	if (TheXDisplay) {
// 		XSync(TheXDisplay, False);
// 		XCloseDisplay(TheXDisplay);
// 		TheXDisplay = NULL;
// 	}
// }
type Color struct { 
    Red uint
    Green uint
    Blue uint
}
func GetPixelColor (d  XDisplay, x, y int) Color {
	red := C.uint(0)
	green := C.uint(0)
	blue := C.uint(0)
	C.get_pixel_color(d, C.int(x), C.int(y), &red, &green, &blue)
	return Color{uint(red),uint(green),uint(blue)};
}

/***********************************************************
Copyright 1987, 1994, 1998  The Open Group

Permission to use, copy, modify, distribute, and sell this software and its
documentation for any purpose is hereby granted without fee, provided that
the above copyright notice appear in all copies and that both that
copyright notice and this permission notice appear in supporting
documentation.

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE OPEN GROUP BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

Except as contained in this notice, the name of The Open Group shall
not be used in advertising or otherwise to promote the sale, use or
other dealings in this Software without prior written authorization
from The Open Group.


Copyright 1987 by Digital Equipment Corporation, Maynard, Massachusetts

                        All Rights Reserved

Permission to use, copy, modify, and distribute this software and its
documentation for any purpose and without fee is hereby granted,
provided that the above copyright notice appear in all copies and that
both that copyright notice and this permission notice appear in
supporting documentation, and that the name of Digital not be
used in advertising or publicity pertaining to distribution of the
software without specific, written prior permission.

DIGITAL DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS SOFTWARE, INCLUDING
ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS, IN NO EVENT SHALL
DIGITAL BE LIABLE FOR ANY SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR
ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS,
WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION,
ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS
SOFTWARE.

******************************************************************/





var XK_leftarrow     =                C.KeySym(0x08fb)  /* U+2190 LEFTWARDS ARROW */
var XK_uparrow       =                C.KeySym(0x08fc)  /* U+2191 UPWARDS ARROW */
var XK_rightarrow    =                C.KeySym(0x08fd)  /* U+2192 RIGHTWARDS ARROW */
var XK_downarrow     =                C.KeySym(0x08fe)  /* U+2193 DOWNWARDS ARROW */

var XK_A             = C.KeySym(                 0x0041 ) /* U+0041 LATIN CAPITAL LETTER A */
var XK_B             = C.KeySym(                 0x0042 ) /* U+0042 LATIN CAPITAL LETTER B */
var XK_C             = C.KeySym(                 0x0043 ) /* U+0043 LATIN CAPITAL LETTER C */
var XK_D             = C.KeySym(                 0x0044 ) /* U+0044 LATIN CAPITAL LETTER D */
var XK_E             = C.KeySym(                 0x0045 ) /* U+0045 LATIN CAPITAL LETTER E */
var XK_F             = C.KeySym(                 0x0046 ) /* U+0046 LATIN CAPITAL LETTER F */
var XK_G             = C.KeySym(                 0x0047 ) /* U+0047 LATIN CAPITAL LETTER G */
var XK_H             = C.KeySym(                 0x0048 ) /* U+0048 LATIN CAPITAL LETTER H */
var XK_I             = C.KeySym(                 0x0049 ) /* U+0049 LATIN CAPITAL LETTER I */
var XK_J             = C.KeySym(                 0x004a ) /* U+004A LATIN CAPITAL LETTER J */
var XK_K             = C.KeySym(                 0x004b ) /* U+004B LATIN CAPITAL LETTER K */
var XK_L             = C.KeySym(                 0x004c ) /* U+004C LATIN CAPITAL LETTER L */
var XK_M             = C.KeySym(                 0x004d ) /* U+004D LATIN CAPITAL LETTER M */
var XK_N             = C.KeySym(                 0x004e ) /* U+004E LATIN CAPITAL LETTER N */
var XK_O             = C.KeySym(                 0x004f ) /* U+004F LATIN CAPITAL LETTER O */
var XK_P             = C.KeySym(                 0x0050 ) /* U+0050 LATIN CAPITAL LETTER P */
var XK_Q             = C.KeySym(                 0x0051 ) /* U+0051 LATIN CAPITAL LETTER Q */
var XK_R             = C.KeySym(                 0x0052 ) /* U+0052 LATIN CAPITAL LETTER R */
var XK_S             = C.KeySym(                 0x0053 ) /* U+0053 LATIN CAPITAL LETTER S */
var XK_T             = C.KeySym(                 0x0054 ) /* U+0054 LATIN CAPITAL LETTER T */
var XK_U             = C.KeySym(                 0x0055 ) /* U+0055 LATIN CAPITAL LETTER U */
var XK_V             = C.KeySym(                 0x0056 ) /* U+0056 LATIN CAPITAL LETTER V */
var XK_W             = C.KeySym(                 0x0057 ) /* U+0057 LATIN CAPITAL LETTER W */
var XK_X             = C.KeySym(                 0x0058 ) /* U+0058 LATIN CAPITAL LETTER X */
var XK_Y             = C.KeySym(                 0x0059 ) /* U+0059 LATIN CAPITAL LETTER Y */
var XK_Z             = C.KeySym(                 0x005a ) /* U+005A LATIN CAPITAL LETTER Z */

var  XK_Home      = C.KeySym(  0xff50) 
var  XK_Left      = C.KeySym(  0xff51) /* Move left, left arrow */
var  XK_Up        = C.KeySym(  0xff52) /* Move up, up arrow */
var  XK_Right     = C.KeySym(  0xff53) /* Move right, right arrow */
var  XK_Down      = C.KeySym(  0xff54) /* Move down, down arrow */
var  XK_Prior     = C.KeySym(  0xff55) /* Prior, previous */
var  XK_Page_Up   = C.KeySym(  0xff55) 
var  XK_Next      = C.KeySym(  0xff56) /* Next */
var  XK_Page_Down = C.KeySym(  0xff56) 
var  XK_End       = C.KeySym(  0xff57) /* EOL */
var  XK_Begin     = C.KeySym(  0xff58) /* BOL */
