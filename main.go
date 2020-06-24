package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/viile/tetris/tetris"
	"math/rand"
	"time"
)

func inputFromTermbox(g *tetris.Game) (err error) {
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				g.Stop()
				return
			case termbox.KeyArrowUp, termbox.KeyArrowDown, termbox.KeyArrowLeft, termbox.KeyArrowRight:
				g.Input(int(ev.Key))
			default:
			}
		}
	}
}

func inputFromDebug(g *tetris.Game) {
	for _ = range time.NewTicker(time.Second * 1).C {
		g.Input(65517)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	g := tetris.NewGame()
	go g.Run()

	if err := inputFromTermbox(g); err != nil {
		inputFromDebug(g)
	}
}
