package main

/*
struct A {
	int i;
	float f;
};

struct B {
	int type; //下划线访问
};

struct C {
	int type; // 是go语言关键字
	float _type; // 屏蔽cgo对type的访问
};

struct D {
	int size: 10;
	float arr[];
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)

	var b C.struct_B
	fmt.Println(b._type)

	var c C.struct_C
	fmt.Println(c._type)


	//var d C.struct_D
	// error
	//fmt.Println(d.size)
	//fmt.Println(a.arr)
}