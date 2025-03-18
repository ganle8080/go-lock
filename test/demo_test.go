package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestXxx(t *testing.T) {
	fmt.Println("helloworld")
	wg.Add(2)
	go Add()
	go Add()
	wg.Wait()
	fmt.Printf("demo.Num: %v\n", demo.Num)
}

type Demo struct {
	Num   int
	Mutex sync.Mutex
}

var demo Demo

var wg sync.WaitGroup

func Add() {
	for i := 0; i < 5000; i++ {
		demo.Mutex.Lock()
		demo.Num++
		demo.Mutex.Unlock()
	}
	wg.Done()
}
