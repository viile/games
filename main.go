package main

import (
	"errors"
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/viile/games/common"
	"github.com/viile/games/snake"
	"github.com/viile/games/tetris"
	"math/rand"
	"time"
)

func inputFromTermbox(m *common.Manager) (err error) {
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
					m.Input(common.DirectUp)
				case 's','S':
					m.Input(common.DirectDown)
				case 'a','A':
					m.Input(common.DirectLeft)
				case 'd','D':
					m.Input(common.DirectRight)
				}
			case termbox.KeyEsc:
				m.Stop()
				return
			case termbox.KeyCtrlC:
				err = errors.New("CtrlC exit game")
				return
			case termbox.KeyArrowUp,termbox.KeyCtrlW:
				m.Input(common.DirectUp)
			case termbox.KeyArrowDown,termbox.KeyCtrlS:
				m.Input(common.DirectDown)
			case termbox.KeyArrowLeft,termbox.KeyCtrlA:
				m.Input(common.DirectLeft)
			case termbox.KeyArrowRight,termbox.KeyCtrlD:
				m.Input(common.DirectRight)
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
			println(r)
		}
	}()
	for {
		m := common.NewManager()
		var g common.Game
		var i int
		println("\033cTerminal Games \n enter '1' start tetris \n enter '2' start snake")
		_,err := fmt.Scanln(&i)
		if err != nil {
			continue
		}
		switch i {
		case 1:
			g = tetris.NewGame()
		case 2:
			g = snake.NewGame()
		default:
			return
		}

		go m.Run(g)

		if err := inputFromTermbox(m);err != nil {
			return
		}
	}

}
