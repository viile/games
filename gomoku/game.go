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

	aiCh chan int
}

func NewGame(v int) *Game {
	g := &Game{
		G: common.NewG(15, 15),
	}
	g.currUserRole = v
	g.status = StatusWhiteWait
	g.currPoint = g.CenterPoint()
	g.Set(g.currPoint,PointCurr{})
	g.aiCh = make(chan int,8)
	go g.ai()
	if !g.isUserFirst() {
		g.aiCh <- 1
	}
	return g
}

func (g *Game) isUserFirst() bool{
	return g.currUserRole == PointWhiteValue
}

func (g *Game) swapStatus() {
	if g.status == StatusWhiteWait {
		g.status = StatusBlackWait
		return
	}

	g.status = StatusWhiteWait
	return
}

func (g *Game) space(p common.Pos) {
	defer g.swapStatus()

	if g.status == StatusWhiteWait {
		g.Set(p, PointWhite{})
	}else {
		g.Set(p, PointBlack{})
	}

	if g.win(PointWhiteValue) || g.win(PointBlackValue) {
		g.Stop()
		return
	}
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


func (g *Game) spacePoint() {
	if (g.currUserRole == PointBlackValue && g.status != StatusBlackWait) ||
		(g.currUserRole == PointWhiteValue && g.status != StatusWhiteWait){
		return
	}
	if g.Get(g.currPoint).Value() > PointCurrValue {
		return
	}

	g.space(g.currPoint)

	g.aiCh <- 1

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
	for _ = range g.aiCh{
		unlock := g.Lock()
		for {
			p := g.RandPoint()
			if g.Get(p).Value() > PointCurrValue {
				continue
			}
			g.space(p)
			break
		}
		unlock()
	}
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
