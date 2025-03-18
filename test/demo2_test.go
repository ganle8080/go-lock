package test

import "sync"

type gobalMapData struct {
	gMutex sync.RWMutex
	Name   string
	maps   map[string]MapData
}

type MapData struct {
	mMutex sync.RWMutex
	Name   string
	maps   map[string]string
}

// 我将创建两个协程，一个在gobalMapData中修改数据，一个在MapData中修改数据
