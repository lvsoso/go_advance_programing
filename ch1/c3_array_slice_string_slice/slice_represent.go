package main

import "fmt"

var (
	a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
	d = c[:2]             // 有2个元素的切片, len为2, cap为3
	e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
	f = c[:0]             // 有0个元素的切片, len为0, cap为3
	g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
	h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
	i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
)

func main() {
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
}
