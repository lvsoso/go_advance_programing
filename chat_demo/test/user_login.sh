#! /bin/sh

curl -v -i -X POST \
  http://127.0.0.1:8888/user/login \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
  "username":"lvsoso",
  "password":"111111"
}'