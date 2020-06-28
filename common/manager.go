package common

import (
	"time"
)

type Manager struct {
	// 每帧间隔,建议50ms
	heartbeat time.Duration
	// 按键输入事件
	inputChan chan int
	// 心跳事件
	hbChan chan int
	// 停止事件
	stopChan chan int
}

func NewManager() *Manager{
	return &Manager{
		heartbeat:time.Millisecond * 50,
		inputChan: make(chan int),
		hbChan:    make(chan int),
		stopChan:  make(chan int,8),
	}
}

func (m *Manager) hbSender() {
	for _ = range time.NewTicker(m.heartbeat).C {
		m.hbChan <- 1
	}
}


func (m *Manager) Stop() {
	m.stopChan <- 1
}

func (m *Manager) Input(i int) {
	m.inputChan <- i
}

func (m *Manager) Run(g Game) {
	go m.hbSender()
	for {
		select {
		case <-m.stopChan:
			g.Stop()
			return
		case i := <-m.inputChan:
			g.InputEvent(i)
		case <-m.hbChan:
			g.HeartbeatEvent()
		}
	}

}
