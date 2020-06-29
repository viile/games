package block

const (
	BlockLUp = iota
	BlockLDown = iota
	BlockLLeft = iota
	BlockLRight = iota
)

type BlockL struct {
	Te
}
//     ▓▓   ▓▓▓▓  ▓▓▓▓▓▓   ▓▓
// ▓▓▓▓▓▓     ▓▓  ▓▓       ▓▓
//            ▓▓           ▓▓▓▓
func NewBlockL(w,h int) *BlockL {
	return &BlockL{Te{Blocks{
		NewBlock(w, h, true),
		NewBlock(w , h-1, false),
		NewBlock(w - 1, h-1, false),
		NewBlock(w - 2, h-1, false)}}}
}

func (b *BlockL) Value() int {
	return BlockLValue
}
func (b *BlockL) Render() string {
	return "🟨"
}

func (b *BlockL) status() int {
	if b.isu() {
		return BlockLUp
	} else if b.isd() {
		return BlockLDown
	} else if b.isl() {
		return BlockLLeft
	} else {
		return BlockLRight
	}

}

func (b *BlockL) isu() bool {
	o := b.Origin()
	for _,v := range b.Blocks {
		if o.X == v.X && o.Y == v.Y {
			continue
		}
		if v.Y >= o.Y {
			return false
		}
	}

	return true
}

func (b *BlockL) isd() bool {
	o := b.Origin()
	for _,v := range b.Blocks {
		if o.X == v.X && o.Y == v.Y {
			continue
		}
		if v.Y <= o.Y {
			return false
		}
	}

	return true
}

func (b *BlockL) isl() bool {
	o := b.Origin()
	for _,v := range b.Blocks {
		if o.X == v.X && o.Y == v.Y {
			continue
		}
		if v.X <= o.X {
			return false
		}
	}

	return true
}

func (b *BlockL) br() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X-1,o.Y,false)
	bps[2] = NewBlock(o.X-1,o.Y+1,false)
	bps[3] = NewBlock(o.X-1,o.Y+2,false)
	return bps
}

func (b *BlockL) bd() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X,o.Y+1,false)
	bps[2] = NewBlock(o.X+1,o.Y+1,false)
	bps[3] = NewBlock(o.X+2,o.Y+1,false)
	return bps
}

func (b *BlockL) bl() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X+1,o.Y,false)
	bps[2] = NewBlock(o.X+1,o.Y-1,false)
	bps[3] = NewBlock(o.X+1,o.Y-2,false)
	return bps
}

func (b *BlockL) bu() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X,o.Y-1,false)
	bps[2] = NewBlock(o.X-1,o.Y-1,false)
	bps[3] = NewBlock(o.X-2,o.Y-1,false)
	return bps
}

func (b *BlockL) Rotate() Blocks {
	switch b.status() {
	case BlockLUp:
		return b.br()
	case BlockLRight:
		return b.bd()
	case BlockLDown:
		return b.bl()
	case BlockLLeft:
		return b.bu()
	default:
		return b.Blocks
	}
}