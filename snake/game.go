package snake

import (
	"container/list"
	"math/rand"
	"sync"
	"time"
)

const (
	PointZero = iota
	PointSnake
	PointFruit
)

const (
	DirectUp = iota
	DirectDown
	DirectLeft
	DirectRight
)

type Pos struct {
	X, Y int
}

type Game struct {
	// y
	height int
	// x
	weight int
	//
	container map[int][]int
	//
	snake *list.List
	//
	currPoint Pos
	//
	currDirect int
	//
	counter int
	//
	inputChan chan int
	//
	hbChan chan int
	//
	stopChan chan int
	// locker
	locker sync.Mutex
}

func NewGame() *Game {
	var weight, height = 16, 12
	g := &Game{
		height:    height,
		weight:    weight,
		container: make(map[int][]int, height),
		snake: list.New(),
		inputChan: make(chan int),
		hbChan:    make(chan int),
		stopChan:  make(chan int,8),
		locker:    sync.Mutex{},
	}
	g.initContainer()
	g.newSnake()
	g.newFruit()
	return g
}

func (g *Game) initContainer()  {
	for w := 0; w < g.weight; w++ {
		g.container[w] = make([]int, g.height)
		for h := 0; h < g.height; h++ {
			g.container[w][h] = PointZero
		}
	}
}

func (g *Game) newSnake()  {
	var w,h = g.weight / 2,g.height / 2
	g.snake.PushBack(Pos{w+1,h})
	g.snake.PushFront(Pos{w,h})
	g.currDirect = DirectLeft
	for e := g.snake.Front(); e != nil; e = e.Next() {
		p := e.Value.(Pos)
		g.container[p.X][p.Y] = PointSnake
	}
}

func (g *Game) newFruit()  {
	var w, h = rand.Intn(g.weight), rand.Intn(g.height)
	if g.container[w][h] == PointZero {
		g.container[w][h] = PointFruit
		return
	}
	for m:=0;m<g.weight;m++ {
		for n := 0; n < g.height; n++ {
			if g.container[w][h] == PointZero {
				g.container[w][h] = PointFruit
				return
			}
		}
	}

	g.stopChan <- 1
}
func (g *Game) up()  {
	defer g.lock()()
	if g.currDirect != DirectDown {
		g.currDirect = DirectUp
	}
	g.move()
}
func (g *Game) down()  {
	defer g.lock()()
	if g.currDirect != DirectUp {
		g.currDirect = DirectDown
	}
	g.move()
}
func (g *Game) left()  {
	defer g.lock()()
	if g.currDirect != DirectRight {
		g.currDirect = DirectLeft
	}
	g.move()
}
func (g *Game) right()  {
	defer g.lock()()
	if g.currDirect != DirectLeft {
		g.currDirect = DirectRight
	}
	g.move()
}
func (g *Game) move() {
	header := g.snake.Front().Value.(Pos)
	tail := g.snake.Back()
	var p Pos
	switch g.currDirect {
	case DirectUp:
		p = Pos{header.X,header.Y + 1}
	case DirectDown:
		p = Pos{header.X,header.Y - 1}
	case DirectLeft:
		p = Pos{header.X - 1,header.Y }
	case DirectRight:
		p = Pos{header.X + 1,header.Y }
	}

	if p.X < 0 || p.X >= g.weight || p.Y < 0 || p.Y >= g.height {
		g.Stop()
		return
	}

	if g.container[p.X][p.Y] == PointSnake {
		g.Stop()
		return
	}

	if g.container[p.X][p.Y] == PointZero {
		g.snake.PushFront(p)
		g.snake.Remove(tail)
		g.container[p.X][p.Y] = PointSnake
		t := tail.Value.(Pos)
		g.container[t.X][t.Y] = PointZero
		return
	}

	if g.container[p.X][p.Y] == PointFruit {
		g.snake.PushFront(p)
		g.container[p.X][p.Y] = PointSnake
		g.newFruit()
		return
	}

	return
}

func (g *Game)Stop(){
	g.stopChan <- 1
}
func (g *Game)Input(i int){
	g.inputChan <- i
}

func (g *Game) Run() {
	go g.hbSender()
	for {
		select {
		case <-g.stopChan:
			println("stop..")
			return
		case i := <-g.inputChan:
			// 移动当前方块位置
			switch i {
			case 65517:
				g.up()
			case 65516:
				g.down()
			case 65515:
				g.left()
			case 65514:
				g.right()
			}
		case <-g.hbChan:
			g.HeartbeatEvent()
		}
	}

}

func (g *Game) lock() func() {
	g.locker.Lock()
	return func() {
		g.locker.Unlock()
	}
}

// 移动位置,计算积分
// 刷新屏幕
func (g *Game) HeartbeatEvent() {
	defer g.lock()()
	g.counter++
	// 每24帧,移动一格
	if g.counter%24 == 0 {
		g.move()
	}
	g.display()
}

func (g *Game) display() {
	display(g.weight, g.height, g.container)
}

func (g *Game) hbSender() {
	for _ = range time.NewTicker(time.Millisecond * 50).C {
		g.hbChan <- 1
	}
}
