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
	"bufio"
	"bytes"
	"unicode/utf8"

	"github.com/gdamore/tcell"
)

// DrawWrappedText will print text to the screen, wrapping where possible to stay inside the bounds supplied
// in the paramters, and without overstepping the bounds of the screen.
func DrawWrappedText(screen ScreenBridge, minX, minY, maxX, maxY int, text string, style tcell.Style) {

	screenX, screenY := screen.Size()

	if maxX > screenX {
		maxX = screenX
	}

	if maxY > screenY {
		maxY = screenY
	}

	x := minX
	y := minY
	textBuffer := bytes.NewBufferString(text)
	lineScanner := bufio.NewScanner(textBuffer)

	lineScanner.Split(bufio.ScanLines)

	for lineScanner.Scan() {

		if y > maxY {
			break
		}

		line := lineScanner.Text()
		lastSpaceIndex := -1
		lastSpaceScreen := -1
		advance := 1
		r := ' '
		runes := []rune(line)

		if len(line) == 0 {
			// advance y, reset x, keep reading
			y++
			x = minX
			continue
		}

		for i := 0; i < len(runes); i++ {

			if y > maxY {
				break
			}

			r = runes[i]
			advance = getAdvance(r)

			if r == ' ' {
				lastSpaceIndex = i
				lastSpaceScreen = x
			}

			if x+advance <= maxX {
				screen.SetContent(x, y, r, nil, style)
				x += advance

			} else {

				if r == ' ' {
					// advance y, reset x
					y++
					x = minX

				} else if lastSpaceIndex > -1 {
					// paint over everything since last space
					Paint(screen, lastSpaceScreen, y, x, y, ' ', style)

					// advance y, reset x, rewind i to last space
					y++
					x = minX
					i = lastSpaceIndex

				} else {
					// advance y, reset x, rewind i by 1
					y++
					x = minX
					i--
				}
			}

			// // if end of string
			if i+1 == len(runes) {
				y++
			}
		}
	}
}

func getAdvance(r rune) int {
	if utf8.RuneLen(r) > 1 {
		return 2
	}

	return 1
}
