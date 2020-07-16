package gomoku

import (
	"github.com/magiconair/properties/assert"
	"github.com/viile/games/common"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := NewGame()

	assert.Equal(t,g.Weight(),g.Height(),"new game")
}

func TestGameWinPoint(t *testing.T) {
	g := NewGame()
	g.Set(common.Pos{0,1},PointBlack{})
	g.Set(common.Pos{0,2},PointBlack{})
	g.Set(common.Pos{0,3},PointWhite{})
	g.Set(common.Pos{0,4},PointBlack{})
	g.Set(common.Pos{0,5},PointBlack{})

	assert.Equal(t,g.winPoint(common.Pos{0,1},common.DirectionTop,PointBlackValue),1,"1 win point")
	assert.Equal(t,g.winPoint(common.Pos{0,2},common.DirectionTop,PointBlackValue),0,"0 win point")
}

func TestGameWinLine(t *testing.T) {
	g := NewGame()
	g.Set(common.Pos{0,1},PointBlack{})
	g.Set(common.Pos{0,2},PointBlack{})
	g.Set(common.Pos{0,3},PointBlack{})
	g.Set(common.Pos{0,4},PointBlack{})
	g.Set(common.Pos{0,5},PointBlack{})

	assert.Equal(t,g.winLine(common.Pos{0,1},LineVertical,PointBlackValue),true,"win line")
	assert.Equal(t,g.winLine(common.Pos{0,2},LineVertical,PointBlackValue),true,"win line")
	assert.Equal(t,g.winLine(common.Pos{0,3},LineVertical,PointBlackValue),true,"win line")
	assert.Equal(t,g.winLine(common.Pos{0,4},LineVertical,PointBlackValue),true,"win line")
	assert.Equal(t,g.winLine(common.Pos{0,5},LineVertical,PointBlackValue),true,"win line")
}

func TestGameWinCheck(t *testing.T) {
	g := NewGame()
	g.Set(common.Pos{0,1},PointBlack{})
	g.Set(common.Pos{0,2},PointBlack{})
	g.Set(common.Pos{0,3},PointBlack{})
	g.Set(common.Pos{0,4},PointBlack{})
	g.Set(common.Pos{0,5},PointBlack{})

	assert.Equal(t,g.winCheck(common.Pos{0,1},PointBlackValue),true,"black win")
	assert.Equal(t,g.win(PointBlackValue),true,"black win")
}
