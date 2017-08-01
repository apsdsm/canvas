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

package canvas

type cell struct {
	Rune rune
}

// A Layer is a drawable area in a canvas that carries its own origin coords
// and max boundaries.
type Layer struct {
	Grid                            [][]cell
	Width, Height, X, Y, MaxX, MaxY int
}

// NewLayer creates a new layer with the specified width and height with the
// spefified offset.
func NewLayer(width, height, x, y int) *Layer {
	l := Layer{}

	l.Grid = make([][]cell, width)

	for i := range l.Grid {
		l.Grid[i] = make([]cell, height)
	}

	l.Width = width
	l.Height = height
	l.X = x
	l.Y = y
	l.MaxX = width - 1
	l.MaxY = height - 1

	return &l
}

// SetRune will set the cell at the given coordinates with the provided rune.
func (l *Layer) SetRune(x, y int, r rune) {
	if x <= l.MaxX && y <= l.MaxY {
		l.Grid[x][y].Rune = r
	}
}

// Normalize will take a set of rect coords and normalize them to be legal layer coordinates.
func (l *Layer) Normalize(minX, minY, maxX, maxY int) (nMinX, nMinY, nMaxX, nMaxY int) {
	nMinX = normalize(minX, 0, l.MaxX)
	nMaxX = normalize(maxX, 0, l.MaxX)
	nMinY = normalize(minY, 0, l.MaxY)
	nMaxY = normalize(maxY, 0, l.MaxY)

	return
}

// normalize a number so that it is inside a range of a min and max.
func normalize(target, min, max int) int {
	if target < min {
		return min
	} else if target > max {
		return max
	} else {
		return target
	}
}
