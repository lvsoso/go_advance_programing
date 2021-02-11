package main

/*
#include <stdint.h>

union B1 {
	int i;
	float f;
};

union B2 {
	int8_t i8;
	int64_t i64;
};

union B {
	int i;
	float f;
};
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var b1 C.union_B1
	fmt.Printf("%T\n", b1)
	var b2 C.union_B2
	fmt.Printf("%T\n", b2)

	var b C.union_B;
	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))
}
