package car

import (
	"github.com/viile/games/common"
)

type Game struct {
	*common.G
	//
	currPoint common.Pos
	//
	lastMoveCounter int
}

func NewGame() *Game {
	g := &Game{
		G:common.NewG(24,8),
	}

	return g
}

