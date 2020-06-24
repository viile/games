package block

type BlockI struct {
	Te
}

func NewBlockI(w,h int) *BlockI {
	return &BlockI{Te{[]Block{
		Block{w, h, true},
		{w + 1, h, false},
		{w + 2, h, false},
		{w + 3, h, false}}}}
}

func (b *BlockI) Rotate() Blocks {
	if b.isUpright() {
		return b.sbh()
	}
	return b.hbs()
}

func (b *BlockI) isUpright() bool {
	return b.Blocks[0].X == b.Blocks[1].X
}

//(1,10) (1,9) (1,8) (1,7)  => (1,10) (2,10) (3,10) (4,10)
func (b *BlockI) sbh() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = o
	bps[1] = Block{o.X+1,o.Y,false}
	bps[2] = Block{o.X+2,o.Y,false}
	bps[3] = Block{o.X+3,o.Y,false}
	return bps
}

//(1,10) (2,10) (3,10) (4,10)  => (1,10) (1,9) (1,8) (1,7)
func (b *BlockI) hbs() Blocks {
	o := b.Origin()
	bps := make(Blocks, len(b.Blocks))
	bps[0] = o
	bps[1] = Block{o.X,o.Y-1,false}
	bps[2] = Block{o.X,o.Y-2,false}
	bps[3] = Block{o.X,o.Y-3,false}
	return bps
}
