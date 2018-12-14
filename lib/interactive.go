package lib

import (
	"log"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func drawX(x, y int, str string, fgColor int, bgColor int) {
	runes := []rune(str)

	for _, r := range runes {
		termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorDefault)
		x += runewidth.RuneWidth(r)
	}
}

func drawBox() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	drawX(10, 10, "Hello", 0, 0)
	termbox.Flush()
}

// Interactive is exec interactive mode
func Interactive() {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}

	defer termbox.Close()
	drawBox()
MAINLOOP:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP
			default:
				drawBox()
			}
		default:
			drawBox()
		}
	}
}
