package common

import (
	"sync"
	"time"
)

type Game interface {
	Run()
	Stop()
	Input(i int)
	InputEvent()
	HeartbeatEvent()
}

type G struct {
	// 屏幕高度
	height int
	// 屏幕宽度
	weight int
	// 屏幕容器
	container map[int][]int
	// 帧数总计
	counter int
	// 每帧间隔,建议50ms
	heartbeat time.Duration
	// 按键输入事件
	inputChan chan int
	// 心跳事件
	hbChan chan int
	// 停止事件
	stopChan chan int
	// locker
	locker sync.Mutex
}

func NewG(w,h int,hb time.Duration) *G{
	g := &G{
		height:    w,
		weight:    h,
		heartbeat:hb,
		container: make(map[int][]int, h),
		inputChan: make(chan int),
		hbChan:    make(chan int),
		stopChan:  make(chan int,1),
		locker:    sync.Mutex{},
	}
	g.initContainer()
	return g
}

func (g *G) initContainer()  {
	for w := 0; w < g.weight; w++ {
		g.container[w] = make([]int, g.height)
		for h := 0; h < g.height; h++ {
			g.container[w][h] = 0
		}
	}
}

func (g *G) Stop() {
	g.stopChan <- 1
}

func (g *G) Input(i int) {
	g.inputChan <- i
}

func (g *G) HeartbeatEvent(){
	g.counter++
}

func (g *G) hbSender() {
	for _ = range time.NewTicker(time.Millisecond * g.heartbeat).C {
		g.hbChan <- 1
	}
}

func (g *G) Run() {
	go g.hbSender()
	for {
		select {
		case <-g.stopChan:
			return
		case i := <-g.inputChan:
			g.Input(i)
		case <-g.hbChan:
			g.HeartbeatEvent()
		}
	}

}