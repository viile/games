package block

type BlockJ struct {
	Te
}

func NewBlockJ(w,h int) BlockI {
	return BlockI{Te{[]Block{
		Block{w, h, true},
		{w , h-1, false},
		{w + 1, h-1, false},
		{w + 2, h-1, false}}}}
}