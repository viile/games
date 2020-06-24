package block

const (
	BlockTUp = iota
	BlockTDown = iota
	BlockTLeft = iota
	BlockTRight = iota
)

type BlockT struct {
	Te
}

//   ▓▓     ▓▓     ▓▓▓▓▓▓    ▓▓
// ▓▓▓▓▓▓   ▓▓▓▓     ▓▓    ▓▓▓▓
//          ▓▓               ▓▓
func NewBlockT(w,h int) *BlockT {
	return &BlockT{Te{[]Block{
		Block{w, h, true},
		{w - 1, h-1, false},
		{w , h-1, false},
		{w + 1, h-1, false}}}}
}

func (b *BlockT) status() int {
	if b.isu() {
		return BlockTUp
	} else if b.isd() {
		return BlockTDown
	} else if b.isl() {
		return BlockTLeft
	} else {
		return BlockTRight
	}

}

func (b *BlockT) isu() bool {
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

func (b *BlockT) isd() bool {
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

func (b *BlockT) isl() bool {
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

func (b *BlockT) br() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X-1,o.Y + 1,false}
	bps[2] = Block{o.X-1,o.Y,false}
	bps[3] = Block{o.X-1,o.Y - 1,false}
	return bps
}

func (b *BlockT) bd() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X + 1,o.Y+1,false}
	bps[2] = Block{o.X,o.Y+1,false}
	bps[3] = Block{o.X - 1,o.Y+1,false}
	return bps
}

func (b *BlockT) bl() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X+1,o.Y+1,false}
	bps[2] = Block{o.X+1,o.Y,false}
	bps[3] = Block{o.X+1,o.Y-1,false}
	return bps
}

func (b *BlockT) bu() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X - 1,o.Y-1,false}
	bps[2] = Block{o.X,o.Y-1,false}
	bps[3] = Block{o.X+1,o.Y-1,false}
	return bps
}

func (b *BlockT) Rotate() Blocks {
	switch b.status() {
	case BlockTUp:
		return b.br()
	case BlockTRight:
		return b.bd()
	case BlockTDown:
		return b.bl()
	case BlockTLeft:
		return b.bu()
	default:
		return b.Blocks
	}
}