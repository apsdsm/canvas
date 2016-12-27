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
		style  = tcell.StyleDefault
	)

	BeforeEach(func() {
		screen = fakes.NewScreenBridge(100, 100)
	})

	It("draws a horizonal line", func() {
		DrawHLine(screen, 50, 0, 100, style)

		for i := 0; i < 100; i++ {
			Expect(screen.GetRuneAt(i, 50)).To(Equal('-'))
		}
	})

	It("draws a vertical line", func() {
		DrawVLine(screen, 50, 0, 100, style)

		for i := 0; i < 100; i++ {
			Expect(screen.GetRuneAt(50, i)).To(Equal('|'))
		}
	})
})
