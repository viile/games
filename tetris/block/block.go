package block

import (
	"github.com/viile/games/common"
	"math/rand"
)

const (
	BlockZero = iota
	BlockIValue
	BlockJValue
	BlockLValue
	BlockOValue
	BlockSValue
	BlockTValue
	BlockZValue
)

type Block struct {
	common.Pos
	IsOrigin bool
}

func NewBlock(x, y int,o bool) Block{
	return Block{common.Pos{x, y},o}
}

func (b Block) Move(x, y int) Block {
	return Block{common.Pos{b.X + x, b.Y + y},b.IsOrigin}
}
func (b Block) Set(x, y int) Block {
	return Block{common.Pos{x, y},b.IsOrigin}
}

type Blocks []Block

func (b Blocks) Origin() Block {
	for _,v := range b {
		if v.IsOrigin {
			return v
		}
	}

	return NewBlock(b[0].X,b[0].Y,true)
}

type Tran func(Block) Block

func Handle(s Blocks,t Tran) Blocks {
	bps := make(Blocks, len(s))
	for k, v := range s {
		bps[k] = t(v)
	}
	return bps
}

type Tetris interface {
	Get() Blocks
	//GetOriginPoint() Block
	Set(Blocks)
	Rotate() Blocks
	Left() Blocks
	Right() Blocks
	Down() Blocks
	Render() string
	Value() int
}

type Te struct {
	Blocks
}

func (b *Te) Get() Blocks {
	return b.Blocks
}

func (b *Te) Set(s Blocks) {
	b.Blocks = s
}

func (b *Te) Rotate() Blocks {
	return b.Blocks
}

func (b *Te) Left() Blocks {
	return Handle(b.Get(),func(v Block) Block {
		return v.Move(-1, 0)
	})
}

func (b *Te) Right() Blocks {
	return Handle(b.Get(),func(v Block) Block {
		return v.Move(1, 0)
	})
}

func (b *Te) Down() Blocks {
	return Handle(b.Get(),func(v Block) Block {
		return v.Move(0, -1)
	})
}
func (b *Te) Value() int {
	return 0
}
func (b *Te) Render() string {
	return ""
}

func NewTetrisBlock(w,h int) Tetris {
	var b Tetris
	switch rand.Int31n(7) + 1 {
	case BlockIValue:
		b = NewBlockI(w,h)
	case BlockJValue:
		b = NewBlockJ(w,h)
	case BlockLValue:
		b = NewBlockL(w,h)
	case BlockOValue:
		b = NewBlockO(w,h)
	case BlockSValue:
		b = NewBlockS(w,h)
	case BlockTValue:
		b = NewBlockT(w,h)
	case BlockZValue:
		b = NewBlockZ(w,h)
	default:
		b = NewBlockI(w,h)
	}

	return b
}
