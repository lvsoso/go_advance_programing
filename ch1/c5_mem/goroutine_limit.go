package main

import (
	"sync"
	"time"
)

var limit = make(chan int, 3)
var work = []func(){
	func() { println("1"); time.Sleep(1 * time.Second) },
	func() { println("2"); time.Sleep(1 * time.Second) },
	func() { println("3"); time.Sleep(1 * time.Second) },
	func() { println("4"); time.Sleep(1 * time.Second) },
	func() { println("5"); time.Sleep(1 * time.Second) },
}

func main() {
	wg := sync.WaitGroup{}
	for _, w := range work {
		go func(w func()) {
			wg.Add(1)
			limit <- 1
			w()
			<-limit
			wg.Done()
		}(w)
	}

	wg.Wait()
}