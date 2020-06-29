package block

type BlockZ struct {
	Te
}

// ▓▓▓▓       ▓▓
//   ▓▓▓▓   ▓▓▓▓
//          ▓▓
func NewBlockZ(w,h int) *BlockZ {
	return &BlockZ{Te{Blocks{
		NewBlock(w, h, true),
		NewBlock(w + 1, h, false),
		NewBlock(w + 1, h-1, false),
		NewBlock(w + 2, h-1, false)}}}
}

func (b *BlockZ) Value() int {
	return BlockZValue
}
func (b *BlockZ) Render() string {
	return "🟫"
}

func (b *BlockZ) Rotate() Blocks {
	if b.isUpright() {
		return b.bh()
	}
	return b.bs()
}

func (b *BlockZ) isUpright() bool {
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

func (b *BlockZ) bh() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X+1,o.Y,false)
	bps[2] = NewBlock(o.X+1,o.Y-1,false)
	bps[3] = NewBlock(o.X+2,o.Y-1,false)
	return bps
}

func (b *BlockZ) bs() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = NewBlock(o.X,o.Y,true)
	bps[1] = NewBlock(o.X,o.Y-1,false)
	bps[2] = NewBlock(o.X-1,o.Y-1,false)
	bps[3] = NewBlock(o.X-1,o.Y-2,false)
	return bps
}