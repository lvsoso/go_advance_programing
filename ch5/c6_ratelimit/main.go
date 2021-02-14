package main

import (
	"fmt"
	"time"
)

func main() {
	var fillInterval = time.Millisecond * 10
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)

	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	// takeAvailable := func(block bool) bool {
	// 	var takenResult bool
	// 	if block {
	// 		select {
	// 		case <-tokenBucket:
	// 			takenResult = true
	// 		}
	// 	} else {
	// 		select {
	// 		case <-tokenBucket:
	// 			takenResult = true
	// 		default:
	// 			takenResult = false
	// 		}
	// 	}

		return takenResult
	}

	go fillToken()
	time.Sleep(time.Hour)
}
