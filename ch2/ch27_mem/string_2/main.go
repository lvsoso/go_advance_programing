package main

/*
#include<stdio.h>

void printString(const char* s, int n) {
	int i;
	for(i = 0; i < n; i++) {
		putchar(s[i]);
	}
	putchar('\n');
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func printString(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))
}


// 直接在IDE里面运行可能看不到输出
func main() {
	s := "hello"
	printString(s)
}