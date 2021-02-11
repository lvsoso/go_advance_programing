## cgo 

### 

#### import "C"
```text
1. 要使用CGO特性，需要安装C/C++构建工具链，在macOS和Linux下需要安装GCC，在Windows下需要安装MinGW工具。同时需要保证环境变量CGO_ENABLED被设置为1，这表示CGO是被启用的状态。在本地构建时CGO默认是启用的，在交叉构建时CGO默认是禁止的。
2. import "C"导入语句需要单独占一行，不能与其他包一同import。向C函数传递参数也很简单，直接转换成对应的C语言类型传递就可以。
3. Go是强类型语言，所以CGO中传递的参数类型必须与声明的类型完全一致，而且传递前必须用"C"中的转换函数转换成对应的C类型，不能直接传入Go中类型的变量。
4. 不同的Go语言包引入的虚拟的C包之间的类型是不能通用的。
5. 在Go语言中方法是依附于类型存在的，不同Go包中引入的虚拟的C包的类型是不同的（main.C与cgo_helper.C类型不同），这导致从它们延伸出来的Go类型也是不同的类型（*main.C.char与*cgo_helper.C.char类型不同），这最终导致了上面代码不能正常工作。
```
#### `#cgo`语句
```text
1. 在import "C"语句前的注释中可以通过#cgo语句设置编译阶段和链接阶段的相关参数。编译阶段的参数主要用于定义相关宏和指定头文件检索路径。链接阶段的参数主要是指定库文件检索路径和要链接的库文件。
2. 由于C/C++遗留的问题，C头文件检索目录可以是相对路径，但是库文件检索目录则需要绝对路径。
3. 主要影响CFLAGS、CPPFLAGS、CXXFLAGS、FFLAGS和LDFLAGS几个编译器环境变量。LDFLAGS用于设置链接时的参数，除此之外的几个变量用于改变编译阶段的构建参数（CFLAGS用于针对C语言代码设置编译参数）。
4. CFLAGS对应C语言特有的编译选项，CXXFLAGS对应C++特有的编译选项，CPPFLAGS则对应C和C++共有的编译选项。
5. 在链接阶段，C和C++的链接选项是通用的，因此这个时候已经不再有C和C++语言的区别，它们的目标文件的类型是相同的。
6. CLAGS对应C语言编译参数（以.c为扩展名）、CPPFLAGS对应C/C++代码编译参数（*.c/*.cc/*.cpp/*.cxx）、CXXFLAGS对应纯C++编译参数（*.cc/*.cpp/*.cxx）。
```
ex1: import "C"
> CFLAGS部分，-D部分定义了宏PNG_DEBUG，值为1；-I定义了头文件包含的检索目录。LDFLAGS部分，-L指定了链接时库文件检索目录，-l指定了链接时需要链接png库。
> 因为C/C++遗留的问题，C头文件检索目录可以是相对目录，但是库文件检索目录则需要绝对路径。
```go
// #cgo CFLAGS: -DPNG_DEBUG=1 -I./include
// #cgo LDFLAGS: -L/usr/local/lib -lpng
// #include <png.h>
import "C"
```

ex2: SRCDIR
```go
// #cgo LDFLAGS: -L${SRCDIR}/libs -lfoo
``` 

展开

```go
// #cgo LDFLAGS: -L/go/src/foo/libs -lfoo
```

ex3: 支持条件选择
```go
// #cgo windows CFLAGS: -DX86=1
// #cgo !windows LDFLAGS: -lm
```


ex4: 不同的系统下cgo对应着不同的c代码
```go
package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
	const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
	const char* os = "darwin";
#elif defined(CGO_OS_LINUX)
	const char* os = "linux";
#else
#	error(unknown os)
#endif
*/
import "C"

func main() {
	print(C.GoString(C.os))
}
```


#### build标志
```text
1. build标志是在Go或CGO环境下的C/C++文件开头的一种特殊的注释。
2. 可以通过-tags命令行参数同时指定多个build标志，它们之间用空格分隔。
```
ex1:
```go
// +build debug

package main

var buildMode = "debug"
```
ex2:
```shell
go build -tags="windows debug"
```

ex3: 逗号链接表示AND,空白分割来表示OR。
```go
// +build linux,386 darwin,!cgo
```
#### pkg-config
过PKG_CONFIG_PATH环境变量指定`pkg-config`工具的检索目录
```text
Name: 库的名字
Cflags:-I/usr/local/include
Libs:-L/usr/local/lib –lxxx2
```

### cgo types

#### 

C语言类型               | CGO类型      | Go语言类型
---------------------- | ----------- | ---------
char                   | C.char      | byte
singed char            | C.schar     | int8
unsigned char          | C.uchar     | uint8
short                  | C.short     | int16
unsigned short         | C.ushort     | uint16
int                    | C.int       | int32
unsigned int           | C.uint      | uint32
long                   | C.long      | int32
unsigned long          | C.ulong     | uint32
long long int          | C.longlong  | int64
unsigned long long int | C.ulonglong | uint64
float                  | C.float     | float32
double                 | C.double    | float64
size_t                 | C.size_t    | uint


```c
typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef float GoFloat32;
typedef double GoFloat64;
```

#### Int

C语言类型 | CGO类型     | Go语言类型
-------- | ---------- | ---------
int8_t   | C.int8_t   | int8
uint8_t  | C.uint8_t  | uint8
int16_t  | C.int16_t  | int16
uint16_t | C.uint16_t | uint16
int32_t  | C.int32_t  | int32
uint32_t | C.uint32_t | uint32
int64_t  | C.int64_t  | int64
uint64_t | C.uint64_t | uint64

#### string & slice

```c
typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;
```

Go1.10针对Go字符串增加了一个_GoString_预定义类型,和两个方法
```c
size_t _GoStringLen(_GoString_ s);
const char *_GoStringPtr(_GoString_ s);
```