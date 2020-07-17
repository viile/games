package gomoku

const (
	PointValue = iota
	PointCurrValue
	PointBlackValue
	PointWhiteValue
)
type PointBlack struct {
}

func (p PointBlack) Render() string {
	return "🔴"
}
func (p PointBlack) Value() int {
	return PointBlackValue
}

type PointWhite struct {
}

func (p PointWhite) Render() string {
	return "⚪️"
}
func (p PointWhite) Value() int {
	return PointWhiteValue
}

type PointCurr struct {
}

func (p PointCurr) Render() string {
	return "🔸"
}
func (p PointCurr) Value() int {
	return PointCurrValue
}
