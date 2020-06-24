package block

import "math/rand"

const (
	blockI = iota
	blockJ
	blockL
	blockO
	blockS
	blockT
	blockZ
)

type Block struct {
	X, Y     int
	IsOrigin bool
}

func (b *Block) Move(x, y int) Block {
	return Block{b.X + x, b.Y + y, b.IsOrigin}
}
func (b *Block) Set(x, y int) Block {
	return Block{x, y, b.IsOrigin}
}

type Blocks []Block
type Tran func(Block) Block

func (s Blocks) Handle(t Tran) Blocks {
	bps := make(Blocks, len(s))
	for k, v := range s {
		bps[k] = t(v)
	}
	return bps
}
func (s Blocks) Origin() Block {
	for _, v := range s {
		if v.IsOrigin {
			return v
		}
	}

	return s[0]
}

type Tetris interface {
	Get() Blocks
	Set(Blocks)
	Rotate() Blocks
	Left() Blocks
	Right() Blocks
	Down() Blocks
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
	return b.Handle(func(v Block) Block {
		return v.Move(-1, 0)
	})
}

func (b *Te) Right() Blocks {
	return b.Handle(func(v Block) Block {
		return v.Move(1, 0)
	})
}

func (b *Te) Down() Blocks {
	return b.Handle(func(v Block) Block {
		return v.Move(0, -1)
	})
}

func NewBlock(w,h int) Tetris {
	var b Tetris
	switch rand.Int31n(7) {
	case blockI:
		b = NewBlockI(w,h)
	case blockJ:
		b = NewBlockJ(w,h)
	case blockL:
		b = NewBlockL(w,h)
	case blockO:
		b = NewBlockO(w,h)
	case blockS:
		b = NewBlockS(w,h)
	case blockT:
		b = NewBlockT(w,h)
	case blockZ:
		b = NewBlockZ(w,h)
	default:
		b = NewBlockI(w,h)
	}

	return b
}
