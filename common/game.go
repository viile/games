package common

import (
	"math/rand"
	"sync"
)

type Game interface {
	Run()
	Stop()
	InputEvent(i int)
	HeartbeatEvent()
}

type G struct {
	// 屏幕高度
	height int
	// 屏幕宽度
	weight int
	// 屏幕容器
	container map[int][]Point
	// 帧数总计
	counter int
	//
	status int
	// locker
	locker sync.Mutex
}

func NewG(w,h int) *G{
	g := &G{
		height:    w,
		weight:    h,
		container: make(map[int][]Point, h),
		locker:    sync.Mutex{},
		status: StatusRunning,
	}
	g.initContainer()
	return g
}

func (g *G) initContainer()  {
	for w := 0; w < g.weight; w++ {
		g.container[w] = make([]Point, g.height)
		for h := 0; h < g.height; h++ {
			g.container[w][h] = P{}
		}
	}
}

func (g *G) Weight() int {
	return g.weight
}

func (g *G) Height() int {
	return g.height
}

func (g *G) AddCounter()  {
	g.counter++
}
func (g *G) Counter() int {
	return g.counter
}

func (g *G) CenterPoint() Pos {
	return NewPos(g.weight / 2,g.height / 2)
}

func (g *G) RandPoint() Pos {
	return NewPos(rand.Intn(g.weight), rand.Intn(g.height))
}

func (g *G) Overflow(p Pos) bool {
	if p.X < 0 || p.X >= g.Weight() || p.Y < 0 || p.Y >= g.Height() {
		return true
	}
	return false
}

func (g *G) Get(p Pos) Point {
	return  g.container[p.X][p.Y]
}

func (g *G) Set(p Pos,v Point)   {
	g.container[p.X][p.Y] = v
}

func (g *G) Each(fn func(p Pos) bool) bool {
	for m:=0;m<g.weight;m++ {
		for n := 0; n < g.height; n++ {
			if fn(NewPos(m,n)) {
				return true
			}
		}
	}

	return false
}

func (g *G) Running() bool {
	return g.status == StatusRunning
}

func (g *G) Lock() func() {
	g.locker.Lock()
	return func() {
		g.locker.Unlock()
	}
}

func (g *G) Display() {
	var str string
	str += "\033c"
	//
	str += " "
	for w := 0; w < g.Weight(); w++ {
		str += "--"
	}
	str += "\n"
	//
	var m = g.Height() / 2
	for h := g.Height() - 1; h >= 0; h-- {
		if !g.Running() && h == m {
			str += "        GAMEOVER\n   press <esc> exit game\n"
		}else {
			str += "|"
			for w := 0; w < g.Weight(); w++ {
				str += g.Get(NewPos(w, h)).Render()
			}
			str += "|\n"
		}
	}
	//
	str += " "
	for w := 0; w < g.Weight(); w++ {
		str += "--"
	}
	str += "\n"
	//
	print(str)
}

func (g *G) Run() {
}
func (g *G) Stop() {
	g.status = StatusStop
	g.Display()
}
func (g *G) InputEvent(i int) {
}
func (g *G) HeartbeatEvent(){
}
