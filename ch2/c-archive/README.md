
```shell
go build --buildmode=c-archive  hello.go

vim hello.c

gcc  -o hello_static hello.c hello.a -lpthread
```