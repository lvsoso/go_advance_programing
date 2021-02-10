package main

import (
	"sync"
	"sync/atomic"
)

var totalAto uint64

func workerAto(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&totalAto, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go workerAto(&wg)
	go workerAto(&wg)
	wg.Wait()
}