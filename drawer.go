//    Copyright 2016 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package tcelltools

import "github.com/gdamore/tcell"

// A Drawer draws shapes and text to a tcell screen.
type Drawer interface {
	DrawHLine(y, xFrom, xTo int, style tcell.Style)
	DrawVLine(x, from, to int, style tcell.Style)
	Paint(startX, startY, endX, endY int, char rune, style tcell.Style)
	DrawWrappedText(minX, minY, maxX, maxY int, text string, style tcell.Style)
	DrawText(x, y, maxX int, text string, style tcell.Style)
}

// TcellDrawer implements Drawer for Tcell
type TcellDrawer struct {
	screen ScreenBridge
}

// NewDrawer creates and returns a new drawer object.
func NewDrawer(screen ScreenBridge) *TcellDrawer {
	t := TcellDrawer{
		screen: screen,
	}

	return &t
}
