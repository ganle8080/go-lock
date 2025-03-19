package test

import (
	"fmt"
	"sync"
	"testing"
)

type gobalMapData struct {
	gMutex sync.RWMutex
	Name   string
	maps   map[string]*MapData
}

type MapData struct {
	mMutex sync.RWMutex
	Name   string
	maps   map[string]int
}

var wg1 sync.WaitGroup

func Add1(tm *gobalMapData) {

	for i := 0; i < 5000; i++ {
		tm.maps["mapdata"].mMutex.Lock()
		tm.maps["mapdata"].maps["1"]++
		tm.maps["mapdata"].mMutex.Unlock()
	}
	wg1.Done()
}

func TestXxx2(t *testing.T) {
	gobal := gobalMapData{
		Name: "gobal",
	}

	mapdata := MapData{
		Name: "mapdata",
	}

	maps := map[string]int{
		"1": 0,
	}

	mapdata.maps = maps
	gobal.maps = make(map[string]*MapData)
	gobal.maps["mapdata"] = &mapdata

	wg1.Add(2)
	go Add1(&gobal)
	go Add1(&gobal)
	wg1.Wait()

	fmt.Printf("最终数据: %v\n", gobal.maps["mapdata"].maps["1"])
}
