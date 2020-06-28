package gomoku

import (
	"github.com/viile/games/common"
)

type Game struct {
	*common.G
	//
	currPoint common.Pos
}
