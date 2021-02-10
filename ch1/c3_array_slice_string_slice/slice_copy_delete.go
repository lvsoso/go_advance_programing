package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:0], a[1:]...)
	fmt.Println(a)
	a = append(a[:0], a[3:]...)
	fmt.Println(a)

	a = []int{1, 2, 3, 4, 5}
	a = a[:copy(a, a[1:])]
	fmt.Println(a) // 删除开头1个元素
	a = a[:copy(a, a[:3])]// 删除开头3个元素
	fmt.Println(a)
}
