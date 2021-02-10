package main

import (
	"fmt"
)

func twiceV1(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

func twiceV2(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}

func main() {
	x := []int{1, 2, 3}
	twiceV1(x)
	fmt.Println(x)
	v := IntSliceHeader{
		Data: x,
		Len: len(x),
		Cap: cap(x),
	}
	twiceV2(v)
	fmt.Println(x)
}
