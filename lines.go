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

// DrawHLine draws a horizontal line on the screen.
func DrawHLine(screen ScreenBridge, y, from, to int, style tcell.Style) {
	for i := from; i < to; i++ {
		screen.SetContent(i, y, '-', nil, style)
	}
}

// DrawVLine draws a vertical line on the screen.
func DrawVLine(screen ScreenBridge, x, from, to int, style tcell.Style) {
	for i := from; i < to; i++ {
		screen.SetContent(x, i, '|', nil, style)
	}
}
