package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
		}
	}
}

func main() {
	cancel := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker2(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}