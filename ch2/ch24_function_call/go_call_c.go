package main
/*
#include <errno.h>

static int add(int a, int b){
	return a + b;
}

static int div(int a, int b){
	if(b == 0){
		errno = EINVAL;
		return 0;
	}
	return a/b;
}

static void noreturn(){}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(1, 1))
	fmt.Println(C.div(9, 3))

	v0, err0 := C.div(2, 1)
	fmt.Println(v0, err0)
	v1, err1 := C.div(2, 0)
	fmt.Println(v1, err1)

	v, err := C.noreturn()
	fmt.Println(err)
	fmt.Printf("%#v", v)
}

