# go_advance_programing
go advance programing


### nc
```shell

nc localhost 1234
{"method":"HelloService.Hello","params":["hello"],"id":0}
echo -e '{"method":"HelloService.Hello","params":["hello2222"],"id":3}' | nc localhost 1234
echo -e '{"method":"HelloService.Hello","params":["hello2222"],"id":3}{"method":"HelloService.Hello","params":["hello33"],"id":4}' | nc localhost 1234

// nc -l 2399

curl localhost:1234/jsonrpc --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
```


```shell
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```


### grpcurl
```text
列出本地服务
grpcurl localhost:1234 list

使用tls
grpcurl localhost:1234 list -cert xxxx.cert -key xxx.key

不使用tls
grpcurl localhost:1234 list -plaintext

使用 unix socket
grpcurl localhost:1234 list -unix

查看HelloService服务的方法列表
grpcurl -plaintext localhost:1234 list HelloService.HelloService

查看HelloService服务的方法详细描述
grpcurl -plaintext localhost:1234 describe HelloService.HelloService

查看类型信息
grpcurl -plaintext localhost:1234 describe HelloService.String

调用方法
grpcurl -plaintext -d '{"value": "gopher"}' localhost:1234 HelloService.HelloService/Hello

调用流方法
grpcurl -plaintext -d @ localhost:1234 HelloService.HelloService/Channel
```




