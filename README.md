# go_advance_programing
go advance programing


```shell

nc localhost 1234
{"method":"HelloService.Hello","params":["hello"],"id":0}
echo -e '{"method":"HelloService.Hello","params":["hello2222"],"id":3}' | nc localhost 1234
echo -e '{"method":"HelloService.Hello","params":["hello2222"],"id":3}{"method":"HelloService.Hello","params":["hello33"],"id":4}' | nc localhost 1234

// nc -l 2399
```
