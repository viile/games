package snake

import (
	"container/list"
	"github.com/viile/games/common"
)

type Game struct {
	*common.G
	//
	snake *list.List
	//
	currPoint common.Pos
	//
	currDirect int
	//
	lastMoveCounter int
}

func NewGame() *Game {
	g := &Game{
		G:common.NewG(16,12),
		snake: list.New(),
	}

	g.newSnake()
	g.newFruit()
	return g
}


func (g *Game) newSnake()  {
	var p = g.CenterPoint()
	g.snake.PushBack(common.NewPos(p.X+1,p.Y))
	g.snake.PushFront(common.NewPos(p.X,p.Y))
	g.currDirect = common.Left
	for e := g.snake.Front(); e != nil; e = e.Next() {
		g.Set(e.Value.(common.Pos),PointSnake{})
	}
}

func (g *Game) newFruit()  {
	var p = g.RandPoint()

	if g.Get(p).Value() == PointValue {
		g.Set(p,PointFruit{})
		return
	}

	if g.Each(func(p common.Pos) bool {
		return g.Get(p).Value() == PointValue
	}) {
		g.Set(p,PointFruit{})
		return
	}

	g.Stop()
}

func (g *Game) move() {
	header := g.snake.Front().Value.(common.Pos)
	tail := g.snake.Back()
	var p common.Pos
	switch g.currDirect {
	case common.Up:
		p = common.NewPos(header.X,header.Y + 1)
	case common.Down:
		p = common.NewPos(header.X,header.Y - 1)
	case common.Left:
		p = common.NewPos(header.X - 1,header.Y )
	case common.Right:
		p = common.NewPos(header.X + 1,header.Y )
	}
	if g.Overflow(p) {
		g.Stop()
		return
	}

	switch g.Get(p).Value() {
	case PointSnakeValue:
		g.Stop()
		return
	case PointValue:
		g.snake.PushFront(p)
		g.Set(p, PointSnake{})
		g.snake.Remove(tail)
		g.Set(tail.Value.(common.Pos), common.P{})
	case PointFruitValue:
		g.snake.PushFront(p)
		g.Set(p, PointSnake{})
		g.newFruit()
	}
	return
}

func (g *Game) InputEvent(i int) {
	if !g.Running() {
		return
	}
	defer g.Lock()()
	switch i {
	case common.Up:
		if g.currDirect == common.Down {
			return
		}
	case common.Down:
		if g.currDirect == common.Up {
			return
		}
	case common.Left:
		if g.currDirect == common.Right {
			return
		}
	case common.Right:
		if g.currDirect == common.Left {
			return
		}
	default:
		return
	}
	g.currDirect = i
	g.lastMoveCounter = g.Counter()
	g.move()
}

func (g *Game) HeartbeatEvent() {
	if !g.Running() {
		return
	}
	defer g.Lock()()
	g.AddCounter()
	// 每24帧,移动一格
	if g.Counter() - g.lastMoveCounter > 12 && g.Counter()%24 == 0 {
		g.move()
	}
	// 刷新屏幕
	g.Display()
}

