package tetris


const (
	PointValue = iota
	PointBlockValue
)
type PointBlock struct {
}

func (p PointBlock) Render() string {
	return "⬛️"
}
func (p PointBlock) Value() int {
	return 1
}

