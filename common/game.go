package common

import (
	"sync"
	"time"
)

type Game interface {
	Run()
	Stop()
	Input(i int)
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
	//
	heartbeat int
	// 按键输入
	inputChan chan int
	// 心跳
	hbChan chan int
	// 停止
	stopChan chan int
	// locker
	locker sync.Mutex
}

func NewG(w,h,hb int) *G{
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

func (g *G) hbSender() {
	for _ = range time.NewTicker(time.Millisecond * time.Duration(g.heartbeat)).C {
		g.hbChan <- 1
	}
}

func (g *G) Input(i int) {
	g.inputChan <- i
}

func (g *G) Run() {
	go g.hbSender()
	for {
		select {
		case <-g.stopChan:
			return
		case <-g.inputChan:
		case <-g.hbChan:
			g.counter++
		}
	}

}

func (g *G) Stop() {
	g.stopChan <- 1
}