package gomoku

import (
	"github.com/viile/games/common"
)

const (
	Line = iota
	LinebaiHorizontal
	LineVertical
	LineLeftFalling
	LineRightFalling
)

type Game struct {
	*common.G
	//
	currPoint common.Pos
}

func NewGame() *Game {
	g := &Game{
		G: common.NewG(15, 15),
	}
	return g
}

func (g *Game) winPoint(p common.Pos, d, v int) (result int) {
	var m = 0
	var u = p

	for m <= 4 {
		m++
		u = u.Move(d)
		if g.Overflow(u) || g.Get(u).Value() != v {
			break
		}
		result++
	}

	return
}

func (g *Game) winLine(p common.Pos, l, v int) bool {
	var result int

	switch l {
	case LinebaiHorizontal:
		result = g.winPoint(p, common.DirectionLeft, v) + g.winPoint(p, common.DirectionRight, v)
	case LineVertical:
		result = g.winPoint(p, common.DirectionTop, v) + g.winPoint(p, common.DirectionDown, v)
	case LineLeftFalling:
		result = g.winPoint(p, common.DirectionTopRight, v) + g.winPoint(p, common.DirectionDownRight, v)
	case LineRightFalling:
		result = g.winPoint(p, common.DirectionTopLeft, v) + g.winPoint(p, common.DirectionDownLeft, v)
	}

	if result >= 4 {
		return true
	}

	return false
}

func (g *Game) winCheck(p common.Pos, v int) bool {
	if g.winLine(p, LinebaiHorizontal, v) || g.winLine(p, LineVertical, v) ||
		g.winLine(p, LineLeftFalling, v) || g.winLine(p, LineRightFalling, v) {
		return true
	}

	return false
}

func (g *Game) win(v int) bool {
	return g.Each(func(p common.Pos) bool {
		return g.winCheck(p, v)
	})
}
