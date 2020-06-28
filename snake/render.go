package snake

const (
	PointValue = iota
	PointSnakeValue
	PointFruitValue
)

type PointSnake struct {
}

func (p PointSnake) Render() string {
	return "🐍"
}
func (p PointSnake) Value() int {
	return PointSnakeValue
}

type PointFruit struct {
}

func (p PointFruit) Render() string {
	return "🍎"
}
func (p PointFruit) Value() int {
	return PointFruitValue
}

