```shell
go build --buildmode=c-shared -o libhello.so  hello.go

gcc  -o hello_dynamic hello.c  -L./   -lhello -lpthread
```