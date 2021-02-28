
## 简单的聊天应用练习

### 需求
- 实现一对一
- 群组对话


### 功能

### goctl

```shell
goctl api -o user.api
goctl api go -api user.api -dir .
go run user.go -f etc/user-api.yaml

goctl model mysql datasource -url="myuser:mypasswd@tcp(127.0.0.1:3306)/chat_demo" -table="user" -dir ./model
```

