package tetris

import "time"

type Game struct {
	// y
	height int
	// x
	weight int
	//
	container map[int][]int
	//
	blocks []int
	//
	currScore int
	//
	maxScore int
	//
	status int
	//
	counter int
	//
	inputChan chan int
	//
	hzChan chan int
	//
	stopChan chan int
}

func NewGame() *Game {
	var height,weight = 80,30
	g := &Game{
		height:height,
		weight:weight,
		container:make(map[int][]int,height),
		inputChan:make(chan int,1),
		hzChan:make(chan int,1),
		stopChan:make(chan int,1),
	}
	for h:= 0;h<g.height;h++ {
		g.container[h] = make([]int,weight)
		for w:=0;w<g.weight;w++ {
			g.container[h][w] = 0
		}
	}
	return g
}

func (g *Game) hz() {
	for _ = range time.NewTicker(time.Millisecond * 41).C {
		g.hzChan <- 1
	}
}

func (g *Game) Run() {
	go g.hz()
	for {
		select {
		case <- g.stopChan:
			return
		case <- g.inputChan:
			// 计算位置
			return
		case <- g.hzChan:
			// 刷新屏幕
			return
		}
	}

}

func (g *Game) Input(i int) {
	g.inputChan <- i
}

func (g *Game) Stop() {
	g.stopChan <- 1
}