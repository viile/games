package block

const (
	BlockJUp = iota
	BlockJDown = iota
	BlockJLeft = iota
	BlockJRight = iota
)

type BlockJ struct {
	Te
}
// ▓▓       ▓▓▓▓  ▓▓▓▓▓▓     ▓▓
// ▓▓▓▓▓▓   ▓▓        ▓▓     ▓▓
//          ▓▓             ▓▓▓▓
func NewBlockJ(w,h int) *BlockJ {
	return &BlockJ{Te{[]Block{
		Block{w, h, true},
		{w , h-1, false},
		{w + 1, h-1, false},
		{w + 2, h-1, false}}}}
}

func (b *BlockJ) status() int {
	if b.isu() {
		return BlockJUp
	} else if b.isd() {
		return BlockJDown
	} else if b.isl() {
		return BlockJLeft
	} else {
		return BlockJRight
	}

}

func (b *BlockJ) isu() bool {
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

func (b *BlockJ) isd() bool {
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

func (b *BlockJ) isl() bool {
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

func (b *BlockJ) br() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X-1,o.Y,false}
	bps[2] = Block{o.X-1,o.Y-1,false}
	bps[3] = Block{o.X-1,o.Y-2,false}
	return bps
}

func (b *BlockJ) bd() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X,o.Y+1,false}
	bps[2] = Block{o.X-1,o.Y+1,false}
	bps[3] = Block{o.X-2,o.Y+1,false}
	return bps
}

func (b *BlockJ) bl() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X+1,o.Y,false}
	bps[2] = Block{o.X+1,o.Y+1,false}
	bps[3] = Block{o.X+1,o.Y+2,false}
	return bps
}

func (b *BlockJ) bu() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = Block{o.X,o.Y,true}
	bps[1] = Block{o.X,o.Y-1,false}
	bps[2] = Block{o.X+1,o.Y-1,false}
	bps[3] = Block{o.X+2,o.Y-1,false}
	return bps
}

func (b *BlockJ) Rotate() Blocks {
	switch b.status() {
	case BlockJUp:
		return b.br()
	case BlockJRight:
		return b.bd()
	case BlockJDown:
		return b.bl()
	case BlockJLeft:
		return b.bu()
	default:
		return b.Blocks
	}
}