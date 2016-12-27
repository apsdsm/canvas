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

	It("writes text with wrapping inside a designated area", func() {

		DrawWrappedText(screen, 0, 0, 10, 10, "this is some text", style)

		Expect(screen.GetLine(0, 0, 9)).To(Equal("this is   "))
		Expect(screen.GetLine(1, 0, 8)).To(Equal("some text"))
	})

	It("wraps text when eol start of new word", func() {
		DrawWrappedText(screen, 0, 0, 8, 8, "abc def ghi jkl", style)

		Expect(screen.GetLine(0, 0, 7)).To(Equal("abc def "))
		Expect(screen.GetLine(1, 0, 6)).To(Equal("ghi jkl"))
	})

	It("wraps text when eol is space", func() {
		DrawWrappedText(screen, 0, 0, 7, 7, "abc def ghi jkl", style)

		Expect(screen.GetLine(0, 0, 6)).To(Equal("abc def"))
		Expect(screen.GetLine(1, 0, 6)).To(Equal("ghi jkl"))
	})

	It("wraps very long text", func() {
		DrawWrappedText(screen, 0, 0, 6, 6, "abcdefghijkl", style)

		Expect(screen.GetLine(0, 0, 5)).To(Equal("abcdef"))
		Expect(screen.GetLine(1, 0, 5)).To(Equal("ghijkl"))
	})

	It("does not write beyond specified area", func() {

		beforeWrite := screen.GetLine(1, 0, 5)

		DrawWrappedText(screen, 0, 0, 6, 0, "abcdefghijkl", style)

		Expect(screen.GetLine(0, 0, 5)).To(Equal("abcdef"))
		Expect(screen.GetLine(1, 0, 5)).To(Equal(beforeWrite))
	})

	It("wraps text with line breaks", func() {
		emptyLine := screen.GetLine(0, 0, 3)

		DrawWrappedText(screen, 0, 0, 6, 6, "abcd\n\nefgh\n\nijkl", style)

		Expect(screen.GetLine(0, 0, 3)).To(Equal("abcd"))
		Expect(screen.GetLine(1, 0, 3)).To(Equal(emptyLine))
		Expect(screen.GetLine(2, 0, 3)).To(Equal("efgh"))
		Expect(screen.GetLine(3, 0, 3)).To(Equal(emptyLine))
		Expect(screen.GetLine(4, 0, 3)).To(Equal("ijkl"))
	})
})
