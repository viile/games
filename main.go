package main

import (
	"github.com/nsf/termbox-go"
	"github.com/viile/tetris/tetris"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	g := tetris.NewGame()
	go g.Run()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				g.Stop()
				return
			case termbox.KeyArrowUp,termbox.KeyArrowDown,termbox.KeyArrowLeft,termbox.KeyArrowRight:
				g.Input(int(ev.Key))
			default:
				continue
			}
		}
	}
}