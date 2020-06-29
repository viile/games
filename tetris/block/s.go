package block

type BlockS struct {
	Te
}

//   ▓▓▓▓   ▓▓
// ▓▓▓▓     ▓▓▓▓
//            ▓▓
func NewBlockS(w,h int) *BlockS {
	return &BlockS{Te{Blocks{
		NewBlock(w, h, true),
		NewBlock(w - 1, h, false),
		NewBlock(w - 1, h-1, false),
		NewBlock(w - 2, h-1, false)}}}
}

func (b *BlockS) Value() int {
	return BlockSValue
}
func (b *BlockS) Render() string {
	return "🟦"
}

func (b *BlockS) Rotate() Blocks {
	if b.isUpright() {
		return b.bh()
	}
	return b.bs()
}

func (b *BlockS) isUpright() bool {
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

func (b *BlockS) bh() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X-1,o.Y,false)
	bps[2] = NewBlock(o.X-1,o.Y-1,false)
	bps[3] = NewBlock(o.X-2,o.Y-1,false)
	return bps
}

func (b *BlockS) bs() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X,o.Y-1,false)
	bps[2] = NewBlock(o.X+1,o.Y-1,false)
	bps[3] = NewBlock(o.X+1,o.Y-2,false)
	return bps
}
