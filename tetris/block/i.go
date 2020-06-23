package block

type BlockI struct {
	Te
}

func (b BlockI) Rotate() Blocks {
	if b.isUpright() {
		return b.sbh()
	}
	return b.hbs()
}

func (b BlockI) isUpright() bool {
	return b.Blocks[0].X == b.Blocks[1].X
}

//(1,10) (1,9) (1,8) (1,7)  => (1,10) (2,10) (3,10) (4,10)
func (b BlockI) sbh() Blocks {
	o := b.Origin()
	bps := make(Blocks,len(b.Blocks))
	bps[0] = o
	for i :=1;i<=3;i++ {
		bps[i] = Block{o.X + i,o.Y,false}
	}
	return bps
}

//(1,10) (2,10) (3,10) (4,10)  => (1,10) (1,9) (1,8) (1,7)
func (b BlockI) hbs() Blocks {
	o := b.Origin()
	bps := make(Blocks,len(b.Blocks))
	bps[0] = o
	for i :=1;i<=3;i++ {
		bps[i] = Block{o.X ,o.Y-i,false}
	}
	return bps
}




