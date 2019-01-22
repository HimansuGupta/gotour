package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	i int
	sync.Mutex
}

func main() {
	sc := new(safeCounter)

	for i := 0; i <= 100; i++ {
		fmt.Println("Before Increment", i)
		go sc.Increment()
		fmt.Println("After Increment", i)
		fmt.Println(" ============= Increment", sc.i)
		//go sc.Decrement()
		//fmt.Println("After IncrDecrement", i)
	}
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	sc.i++
	sc.Unlock()
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}
