package block

type BlockL struct {
	Te
}

func NewBlockL(w,h int) BlockL {
	return BlockL{Te{[]Block{
		Block{w, h, true},
		{w , h -1, false},
		{w -1, h-1, false},
		{w -2, h-1, false}}}}
}