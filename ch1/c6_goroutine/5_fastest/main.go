package main

import "fmt"

func main() {
	ch := make(chan string, 32)

	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}

func searchByGoogle(s string) string {
	return s
}

func searchByBing(s string) string {
	return s
}

func searchByBaidu(s string) string {
	return s
}