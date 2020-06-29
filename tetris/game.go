package tetris

import (
	"fmt"
	"github.com/viile/games/common"
	"github.com/viile/games/tetris/block"
)

type Game struct {
	*common.G
	//
	currBlock block.Tetris
	//
	nextBlock int
	//
	currScore int
}

func NewGame() *Game {
	g := &Game{
		G:common.NewG(16,12),
	}
	g.newBlock()
	return g
}

func (g *Game) InputEvent(i int) {
	defer g.Lock()()
	switch i {
	case common.DirectUp:
		g.move(g.currBlock.Rotate)
	case common.DirectDown:
		g.move(g.currBlock.Down)
	case common.DirectLeft:
		g.move(g.currBlock.Left)
	case common.DirectRight:
		g.move(g.currBlock.Right)
	}
}

func (g *Game) cover(o, s block.Blocks) bool {
	for _, v := range s {
		if v.X >= g.Weight() || v.X < 0 {
			return true
		}
		if v.Y >= g.Height() || v.Y < 0 {
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
		if g.Get(common.NewPos(v.X,v.Y)).Value() > 0 {
			return true
		}
	}

	return false
}
func (g *Game) clean(s block.Blocks) {
	for _, v := range s {
		g.Set(common.NewPos(v.X,v.Y),common.P{})
	}
}
func (g *Game) write(s block.Blocks) {
	for _, v := range s {
		g.Set(common.NewPos(v.X,v.Y),g.currBlock)
	}
}
func (g *Game) move(fn func() block.Blocks) bool {
	o := g.currBlock.Get()
	s := fn()
	if g.cover(o, s) {
		return false
	}
	g.currBlock.Set(s)
	g.clean(o)
	g.write(s)
	return true
}


//
func (g *Game) HeartbeatEvent() {
	defer g.Lock()()

	g.AddCounter()
	// 每24帧,移动当前方块往下一格
	if g.Counter()%24 == 0 {
		if !g.move(g.currBlock.Down) {
			// 方块无法继续下降时,进行新方块检测
			if g.checkBlock() {
				// 消行计算
				g.calc()
				g.newBlock()
			}
		}
	}
	// 刷新屏幕
	g.Display()
	g.display()
}

func (g *Game) display() {
	print(fmt.Sprintf("current score %d \n",g.currScore))
}

// 计算是否需要消除x行
func (g *Game) calc() {
	for h := 0; h < g.Height(); h++ {
		fn := func() bool {

			for w := 0; w < g.Weight(); w++ {
				if g.Get(common.NewPos(w,h)).Value() == 0 {
					return false
				}
			}
			return true
		}

		if fn() {
			g.currScore++
			// 消除一行,并下移所有方块
			for hh := h; hh < g.Height(); hh++ {
				for ww := 0; ww < g.Weight(); ww++ {
					if hh +1 < g.Height() {
						g.Set(common.NewPos(ww,hh),g.Get(common.NewPos(ww,hh+1)))
					}else {
						g.Set(common.NewPos(ww,hh),common.P{})
					}
				}
			}


			h--
		}
	}
}

// 产生新的方块
func (g *Game) newBlock(){
	var b = block.NewTetrisBlock(g.Weight() / 2,g.Height() - 1)

	for _, v := range b.Get() {
		if g.Get(v.Pos).Value() > block.BlockZero {
			g.Stop()
			return
		}
	}

	g.currBlock = b
}

// 检查是否需要新的方块 .如果当前方块下方存在方块或边界
func (g *Game) checkBlock() bool{
	c := g.currBlock.Get()
	for _,v := range c {
		// 检测下方是否是边界
		if v.Y <= 0 {
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
		if g.Get(v.Pos).Value() > block.BlockZero {
			return true
		}
	}

	return false
}

