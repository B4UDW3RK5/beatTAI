// Copyright (c) 2014 Peter Hellberg (http://c7.se/)
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// To build this on Debian and Ubuntu, install libgtk-3-dev first.

package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
	"github.com/brandondube/tai"
)

func main() {
	err := ui.Main(func() {
		w := ui.NewWindow("beatTAI GUI", 205, 24, false)
		w.SetMargined(true)

		l := ui.NewLabel("")
		w.SetChild(l)

		w.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		go func() {
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			for now := range ticker.C {
				ui.QueueMain(func() {
					l.SetText(timeInTaiString(now))
				})
			}
		}()

		w.Show()
	})
	if err != nil {
		panic(err)
	}
}

func timeInTaiString(local time.Time) string {
	g := tai.FromTime(local).AsGregorian()
	hourSeconds := g.Hour * 3600
	minuteSeconds := g.Min * 60
	seconds := g.Sec
	beatTai := float64(hourSeconds+minuteSeconds+seconds) / 86.4

	return fmt.Sprintf("@%06.2f (%02d:%02d:%02d)", beatTai, local.Hour(), local.Minute(), local.Second())
}
