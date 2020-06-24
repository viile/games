package snake

import (
	"container/list"
	"sync"
)

type Pos struct {
	X, Y int
}

type Game struct {
	// y
	height int
	// x
	weight int
	//
	container map[int][]int
	//
	snake list.List

	//
	inputChan chan int
	//
	hbChan chan int
	//
	stopChan chan int
	// locker
	locker sync.Mutex
}
