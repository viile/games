package gomoku

import (
	"github.com/viile/games/common"
)

type Game struct {
	*common.G
	//
	currPoint common.Pos
}

func NewGame() *Game {
	g := &Game{
		G:common.NewG(15,15),
	}
	return g
}

func (g *Game) winCheckPoint(p common.Pos) (bool) {
	return false
}


func (g *Game) winCheck(p common.Pos,v int) (bool) {
	return g.Each(g.winCheckPoint)
}