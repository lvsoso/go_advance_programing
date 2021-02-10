package main

import "sync"

var (
	instance1 *singleton
	once sync.Once
)

func Instance1() *singleton {
	once.Do(func() {
		instance1 = &singleton{}
	})
	return instance1
}