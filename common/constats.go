package common

const (
	DirectUp = 65517
	DirectDown = 65516
	DirectLeft = 65515
	DirectRight = 65514
)

const (
	StatusRunning = iota
	StatusStop
)

type Pos struct {
	X,Y int
}

func NewPos(x,y int) Pos{
	return Pos{x,y}
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


