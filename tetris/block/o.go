package block

type BlockO struct {
	Te
}

func NewBlockO(w,h int) *BlockO {
	return &BlockO{Te{[]Block{
		{w, h, true},
		{w + 1, h, false},
		{w , h-1, false},
		{w + 1, h-1, false}}}}
}
