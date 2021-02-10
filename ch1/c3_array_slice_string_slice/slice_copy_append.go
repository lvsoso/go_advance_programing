package main

import "fmt"

func main() {
	a := []int{0, 1, 2}
	a = append(a, 0)
	i := 1
	copy(a[i+1:], a[i:])
	a[i] = 4
	fmt.Println(a)
}
