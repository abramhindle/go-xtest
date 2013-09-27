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

// #cgo pkg-config: x11 xtst
// #include <stdlib.h>
// #include <string.h>
// #include <unistd.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xatom.h>
// #include <X11/Xlocale.h>
// #include <X11/Intrinsic.h>
// #include <X11/StringDefs.h>
// #include <X11/keysym.h>
// #include <X11/extensions/XTest.h>
import "C"

import (
	//"errors"
	//"unsafe"
)


func XOpenDisplay() *C.Display {
	disp := C.XOpenDisplay(nil)
	return disp;
}

func XCloseDisplay(disp *C.Display) {
	C.XSync(disp, 0)
	C.XCloseDisplay(disp);
}

func PressMouseButton(disp *C.Display, button int) C.int {
	retval := C.XTestFakeButtonEvent(disp, C.uint(button), C.int(1), C.ulong(1))
	return retval;
}
func ReleaseMouseButton(disp *C.Display, button int) C.int {
	retval := C.XTestFakeButtonEvent(disp, C.uint(button), C.int(0), C.ulong(1))
	return retval;
}
