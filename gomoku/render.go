package gomoku

const (
	PointValue = iota
	PointBlackValue
	PointWhiteValue
)
type PointBlack struct {
}

func (p PointBlack) Render() string {
	return "️⚫"
}
func (p PointBlack) Value() int {
	return PointBlackValue
}

type PointWhite struct {
}

func (p PointWhite) Render() string {
	return "⚪"
}
func (p PointWhite) Value() int {
	return PointWhiteValue
}
