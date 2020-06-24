package main

import (
	"github.com/nsf/termbox-go"
	"github.com/viile/tetris/tetris"
	"log"
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
			case termbox.KeyArrowUp,termbox.KeyCtrlW:
				g.Input(65517)
			case termbox.KeyArrowDown,termbox.KeyCtrlS:
				g.Input(65516)
			case termbox.KeyArrowLeft,termbox.KeyCtrlA:
				g.Input(65515)
			case termbox.KeyArrowRight,termbox.KeyCtrlD:
				g.Input(65514)
			default:
				continue
			}
		}
	}
}


func main() {
	rand.Seed(time.Now().UnixNano())
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()

	g := tetris.NewGame()
	go g.Run()

	log.Fatalln(inputFromTermbox(g))
}
