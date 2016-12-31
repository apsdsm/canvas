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

package tcelltools_test

import (
	"github.com/apsdsm/binder/fakes"
	. "github.com/apsdsm/tcelltools"
	"github.com/gdamore/tcell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("drawing wrapped text", func() {

	var (
		screen *fakes.ScreenBridge
		style  tcell.Style
		drawer Drawer
	)

	BeforeEach(func() {
		style = tcell.StyleDefault
		screen = fakes.NewScreenBridge(100, 100)
		drawer = NewDrawer(screen)
	})

	It("draws text", func() {

		var drawTextTests = []struct {
			x, y, len int
			text      string
		}{
			{0, 0, 3, "ほげ"},
			{2, 4, 7, "alphabet"},
		}

		for _, test := range drawTextTests {
			drawer.DrawText(test.x, test.y, test.text, style, 20)
			Expect(screen.GetLine(test.y, test.x, test.x+test.len)).To(Equal(test.text))
		}
	})
})
