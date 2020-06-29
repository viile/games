package block

type BlockO struct {
	Te
}

func NewBlockO(w,h int) *BlockO {
	return &BlockO{Te{Blocks{
		NewBlock(w, h, true),
		NewBlock( w+ 1, h, false),
		NewBlock( w, h-1, false),
		NewBlock( w+ 1, h-1, false)}}}
}

func (b *BlockO) Value() int {
	return BlockOValue
}
func (b *BlockO) Render() string {
	return "🟩"
}
