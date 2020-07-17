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

const (
	Status = iota
	StatusWhiteWait
	StatusBlackWait
	StatusStop
)

type Game struct {
	*common.G
	// status
	status int
	//
	currUserRole int
	//
	currPoint common.Pos
}

func NewGame(v int) *Game {
	g := &Game{
		G: common.NewG(15, 15),
	}
	g.currUserRole = v
	g.status = StatusWhiteWait
	g.currPoint = g.CenterPoint()
	g.Set(g.currPoint,PointCurr{})
	if g.currUserRole == PointBlackValue {
		go g.ai()
	}
	return g
}

func (g *Game) moveCurrPoint(d int) {
	var p common.Pos
	switch d {
	case common.Up:
		p = g.currPoint.Top()
	case common.Down:
		p = g.currPoint.Down()
	case common.Left:
		p = g.currPoint.Left()
	case common.Right:
		p = g.currPoint.Right()
	}
	if g.Overflow(p) {
		return
	}
	if g.Get(g.currPoint).Value() == PointCurrValue {
		g.Set(g.currPoint, common.P{})
	}
	g.currPoint = p
	if g.Get(p).Value() <= 0 {
		g.Set(g.currPoint, PointCurr{})
	}
}


func (g *Game) swapStatus() {
	if g.status == StatusWhiteWait {
		g.status = StatusBlackWait
		return
	}

	g.status = StatusWhiteWait
	return
}

func (g *Game) spacePoint() {
	if (g.currUserRole == PointBlackValue && g.status != StatusBlackWait) ||
		(g.currUserRole == PointWhiteValue && g.status != StatusWhiteWait){
		return
	}
	if g.Get(g.currPoint).Value() > PointCurrValue {
		return
	}
	if g.currUserRole == PointBlackValue {
		g.Set(g.currPoint, PointBlack{})
	}else {
		g.Set(g.currPoint, PointWhite{})
	}

	g.swapStatus()

	go g.ai()

	return
}

func (g *Game) InputEvent(i int) {
	defer g.Lock()()
	switch i {
	case common.Up,common.Down,common.Left,common.Right:
		g.moveCurrPoint(i)
	case common.Space:
		g.spacePoint()
	}
}

//
func (g *Game) HeartbeatEvent() {
	defer g.Lock()()

	g.AddCounter()
	// 刷新屏幕
	g.Display()
}

func (g *Game) ai()  {
	defer g.Lock()()
	for {
		p := g.RandPoint()
		if g.Get(p).Value() > PointCurrValue {
			continue
		}
		if g.currUserRole == PointBlackValue {
			g.Set(p, PointWhite{})
		}else {
			g.Set(p, PointBlack{})
		}
		break
	}

	g.swapStatus()
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
	var result int = 1
	if g.Get(p).Value() != v {
		return false
	}
	switch l {
	case LinebaiHorizontal:
		result += g.winPoint(p, common.DirectionLeft, v) + g.winPoint(p, common.DirectionRight, v)
	case LineVertical:
		result += g.winPoint(p, common.DirectionTop, v) + g.winPoint(p, common.DirectionDown, v)
	case LineLeftFalling:
		result += g.winPoint(p, common.DirectionTopRight, v) + g.winPoint(p, common.DirectionDownRight, v)
	case LineRightFalling:
		result += g.winPoint(p, common.DirectionTopLeft, v) + g.winPoint(p, common.DirectionDownLeft, v)
	}

	if result >= 5 {
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
