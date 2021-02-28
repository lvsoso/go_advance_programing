#! /bin/sh

curl -v -i -X POST \
  http://127.0.0.1:8888/user/register \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
        "username":"lvsoso",
        "nickname":"WuErLv",
        "gender":"male",
        "password":"111111",
        "mobile":"13100001111"
}'
