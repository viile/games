package common

const (
	Up    = 65517
	Down  = 65516
	Left  = 65515
	Right = 65514
	Space = 0x20
)

const (
	//
	Direction = iota
	// ↑
	DirectionTop
	// ↓
	DirectionDown
	// ←
	DirectionLeft
	// →
	DirectionRight
	// ↖
	DirectionTopLeft
	// ↗
	DirectionTopRight
	// ↙
	DirectionDownLeft
	// ↘
	DirectionDownRight
)

const (
	StatusRunning = iota
	StatusStop
)

type Pos struct {
	X, Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func (p Pos) Top() Pos {
	return Pos{p.X, p.Y + 1}
}
func (p Pos) Down() Pos {
	return Pos{p.X, p.Y - 1}
}
func (p Pos) Left() Pos {
	return Pos{p.X - 1, p.Y}
}
func (p Pos) Right() Pos {
	return Pos{p.X + 1, p.Y}
}
func (p Pos) TopLeft() Pos {
	return Pos{p.X - 1, p.Y + 1}
}
func (p Pos) TopRight() Pos {
	return Pos{p.X + 1, p.Y + 1}
}
func (p Pos) DownLeft() Pos {
	return Pos{p.X - 1, p.Y - 1}
}
func (p Pos) DownRight() Pos {
	return Pos{p.X + 1, p.Y - 1}
}

func (p Pos) Move(d int) Pos {
	switch d {
	case DirectionTop:
		return p.Top()
	case DirectionDown:
		return p.Down()
	case DirectionLeft:
		return p.Left()
	case DirectionRight:
		return p.Right()
	case DirectionTopLeft:
		return p.TopLeft()
	case DirectionTopRight:
		return p.TopRight()
	case DirectionDownLeft:
		return p.DownLeft()
	case DirectionDownRight:
		return p.DownRight()
	}

	return p
}

type Point interface {
	Render() string
	Value() int
}

type P struct {
}

func (p P) Render() string {
	return "  "
}

func (p P) Value() int {
	return 0
}
