package tetris

import (
	"github.com/viile/tetris/tetris/block"
	"log"
	"sync"
	"time"
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
	currBlock block.Tetris
	//
	nextBlock int
	//
	currScore int
	//
	maxScore int
	//
	status int
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
	var weight, height = 12, 16
	g := &Game{
		height:    height,
		weight:    weight,
		container: make(map[int][]int, height),
		inputChan: make(chan int),
		hbChan:    make(chan int),
		stopChan:  make(chan int),
		locker:    sync.Mutex{},
	}
	g.currBlock = block.BlockI{block.Te{[]block.Block{block.Block{7, 5, true}, {8, 5, false}, {9, 5, false}, {10, 5, false}}}}

	for w := 0; w < g.weight; w++ {
		g.container[w] = make([]int, height)
		for h := 0; h < g.height; h++ {
			g.container[w][h] = 0
		}
	}
	return g
}

func (g *Game) Run() {
	go g.hbSender()
	for {
		select {
		case <-g.stopChan:
			return
		case i := <-g.inputChan:
			// 移动当前方块位置
			switch i {
			case 65517:
				g.rotate()
			case 65516:
				g.down()
			case 65515:
				g.left()
			case 65514:
				g.right()
			}
		case <-g.hbChan:
			// 移动位置,计算积分
			// 刷新屏幕
			g.counter++
			g.hb()
			g.display()
		}
	}

}

func (g *Game) debug() {
	for _, v := range g.container {
		log.Println(v)
	}
}
func (g *Game) display() {
	display(g.weight, g.height, g.container)
}

func (g *Game) lock() func() {
	g.locker.Lock()
	return func() {
		g.locker.Unlock()
	}
}

func (g *Game) cover(o, s block.Blocks) bool {
	for _, v := range s {
		if v.X >= g.weight || v.X < 0 {
			return true
		}
		if v.Y >= g.height || v.Y < 0 {
			return true
		}
		fn := func(v block.Block) bool {
			for _, vv := range o {
				if v.X == vv.X && v.Y == vv.Y {
					return true
				}
			}
			return false
		}
		if fn(v) {
			continue
		}
		if g.container[v.X][v.Y] > 0 {
			return true
		}
	}

	return false
}
func (g *Game) clean(s block.Blocks) {
	for _, v := range s {
		g.container[v.X][v.Y] = 0
	}
}
func (g *Game) write(s block.Blocks) {
	for _, v := range s {
		g.container[v.X][v.Y] = 1
	}
}
func (g *Game) move(fn func() block.Blocks) {
	o := g.currBlock.Get()
	s := fn()
	if g.cover(o, s) {
		return
	}
	g.clean(o)
	g.currBlock = block.BlockI{block.Te{s}}
	g.write(s)
}
func (g *Game) rotate() {
	defer g.lock()()
	g.move(g.currBlock.Rotate)
}
func (g *Game) down() {
	defer g.lock()()
	g.move(g.currBlock.Down)
}
func (g *Game) left() {
	defer g.lock()()
	g.move(g.currBlock.Left)
}
func (g *Game) right() {
	defer g.lock()()
	g.move(g.currBlock.Right)

}
func (g *Game) hb() {
	defer g.lock()()
	g.downBlcok()
	g.newBlock()
	g.calc()
}

// 计算是否需要消除x行
func (g *Game) calc() {

}

// 检查是否需要新的方块 .如果当前方块下方存在方块或边界.则需要新的方块
func (g *Game) newBlock() {

}

// 每24帧,移动当前方块往下一格
func (g *Game) downBlcok() {
	if g.counter%24 == 0 {
		g.move(g.currBlock.Down)
	}
}

func (g *Game) hbSender() {
	for _ = range time.NewTicker(time.Millisecond * 50).C {
		g.hbChan <- 1
	}
}

func (g *Game) Input(i int) {
	g.inputChan <- i
}

func (g *Game) Stop() {
	g.stopChan <- 1
}
