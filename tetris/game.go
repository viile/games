package tetris

import (
	"github.com/viile/tetris/tetris/block"
	"log"
	"math/rand"
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
	g.initContainer()
	g.newBlock()
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

func (g *Game) initContainer()  {
	for w := 0; w < g.weight; w++ {
		g.container[w] = make([]int, g.height)
		for h := 0; h < g.height; h++ {
			g.container[w][h] = 0
		}
	}
}

func (g *Game) debug() {
	for _, v := range g.container {
		log.Println(v)
	}
}
func (g *Game) display() {
	display(g.currScore,g.weight, g.height, g.container)
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
	g.currBlock.Set(s)
	g.clean(o)
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
	if g.checkBlock() {
		g.newBlock()
	}
	g.calc()
}

// 计算是否需要消除x行
func (g *Game) calc() {
	for h := 0; h < g.height; h++ {
		fn := func() bool {

			for w := 0; w < g.weight; w++ {
				if g.container[w][h] == 0 {
					return false
				}
			}
			return true
		}

		if fn() {
			g.currScore++
			// 消除一行,并下移所有方块
			for hh := h; hh < g.height; hh++ {
				for ww := 0; ww < g.weight; ww++ {
					if hh +1 < g.height {
						g.container[ww][hh] = g.container[ww][hh+1]
					}else {
						g.container[ww][hh] = 0
					}
				}
			}


			h--
		}
	}
}

// 产生新的方块
func (g *Game) newBlock(){
	w := g.weight / 2
	h := g.height - 1
	var b block.Tetris
	switch rand.Int31n(7) {
	case BlockI:
		b = block.NewBlockI(w,h)
	case BlockJ:
		b = block.NewBlockJ(w,h)
	case BlockL:
		b = block.NewBlockL(w,h)
	case BlockO:
		b = block.NewBlockO(w,h)
	case BlockS:
		b = block.NewBlockS(w,h)
	case BlockT:
		b = block.NewBlockT(w,h)
	case BlockZ:
		b = block.NewBlockZ(w,h)
	default:
		b = block.NewBlockI(w,h)
	}

	for _, v := range b.Get() {
		if g.container[v.X][v.Y] == 1 {
			g.stopChan <- 1
		}
	}

	g.currBlock = b
}

// 检查是否需要新的方块 .如果当前方块下方存在方块或边界
func (g *Game) checkBlock() bool{
	c := g.currBlock.Get()
	for _,v := range c {
		// 检测下方是否是边界
		if v.Y -1 < 0 {
			return true
		}
		// 检测下方是否是方块内部
		fn := func() bool {
			for _, vv := range c {
				if vv.X == v.X && vv.Y == v.Y - 1 {
					return true
				}
			}
			return false
		}
		if fn() {
			continue
		}
		// 检测下方是否存在其他方块
		if g.container[v.X][v.Y-1] == 1 {
			return true
		}
	}

	return false
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
