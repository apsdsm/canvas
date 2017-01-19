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

import (
	"unicode/utf8"

	"github.com/gdamore/tcell"
)

// DrawText draws text to the x y coords, automatically allocating enough space for double width chars.
func (d *TcellDrawer) DrawText(x, y, maxX int, text string, style tcell.Style) {
	lasti := 0
	used := 0

	// get the length of the string measured in cells. Most runes that have a length of >1 require two
	// cells to be drawn. Everything else fits in 1. It's a fairly primitive appraoch to measuring but
	// it works English/Japanese text. A more reliable approach would not be unwelcome.
	textWidth := 0

	for _, r := range text {
		if utf8.RuneLen(r) > 1 {
			textWidth += 2
		} else {
			textWidth++
		}
	}

	for i, r := range text {

		// double width characters like 「あ」 take up 2 cells when drawn to the terminal, so before we
		// decide the new x position, we check the width of the character. Big characters will usually
		// take up more than 1 byte. ASCII characters are usually just 1 byte long. This code should
		// be refactored into its own method.
		if i > lasti {
			if i-lasti > 1 {
				x += 2
				used += 2
			} else {
				x++
				used++
			}
			lasti = i
		}

		// if the real string width is larger than the max width and we only have 3 spaces left before
		// we hit the max size, fill that space with dots (drawing an ellipse).
		if textWidth > maxX && used == maxX-3 {
			for e := x; e <= x+2; e++ {
				d.screen.SetContent(e, y, '.', nil, style)
			}
			return
		}

		d.screen.SetContent(x, y, r, nil, style)
	}
}
