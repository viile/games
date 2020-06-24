package block

type BlockT struct {
	Te
}

func NewBlockT(w,h int) BlockT {
	return BlockT{Te{[]Block{
		Block{w, h, true},
		{w - 1, h-1, false},
		{w , h-1, false},
		{w + 1, h-1, false}}}}
}