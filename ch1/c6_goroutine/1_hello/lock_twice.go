package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	go func(){
		fmt.Print("你好, 世界")
		mu.Unlock()
	}()
	mu.Lock()
}

