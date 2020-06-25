package main

import (
	"github.com/nsf/termbox-go"
	"github.com/viile/games/tetris"
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
			case 0:
				switch ev.Ch {
				case 'w','W':
					g.InputUp()
				case 's','S':
					g.InputDown()
				case 'a','A':
					g.InputLeft()
				case 'd','D':
					g.InputRight()
				}
			case termbox.KeyEsc:
				g.Stop()
				return
			case termbox.KeyArrowUp,termbox.KeyCtrlW:
				g.InputUp()
			case termbox.KeyArrowDown,termbox.KeyCtrlS:
				g.InputDown()
			case termbox.KeyArrowLeft,termbox.KeyCtrlA:
				g.InputLeft()
			case termbox.KeyArrowRight,termbox.KeyCtrlD:
				g.InputRight()
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
