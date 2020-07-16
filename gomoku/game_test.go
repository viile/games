package gomoku

import (
	"github.com/magiconair/properties/assert"
	"github.com/viile/games/common"
	"testing"
)

func TestGameWin(t *testing.T) {
	g := NewGame()
	g.Set(common.Pos{0,1},PointBlack{})
	g.Set(common.Pos{0,2},PointBlack{})
	g.Set(common.Pos{0,3},PointBlack{})
	g.Set(common.Pos{0,4},PointBlack{})
	g.Set(common.Pos{0,5},PointBlack{})

	assert.Equal(t,g.win(PointBlackValue),true,"black win")
	assert.Equal(t,g.win(PointWhiteValue),false,"white lose")
}
