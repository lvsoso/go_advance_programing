package main

/*
#include <stdlib.h>

void* makeslice(size_t memsize){
	return malloc(memsize);
}
*/
import "C"
import "unsafe"

func makeByteSlice(n int) []byte {
	p := C.makeslice(C.size_t(n))
	// panic: runtime error: slice bounds out of range [::4294967297] with length 2147483648
	return ((*[1<<31]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func main() {
	s := makeByteSlice(1<<32+1)
	s[len(s)-1] = 255
	print(s[len(s)-1])
	freeByteSlice(s)
}