#! /in/sh

curl -X GET \
  http://127.0.0.1:8888/user/info \
  -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ1MTA2NjQsImlhdCI6MTYxNDQyNDI2NCwidXNlcklkIjoxfQ.8Dmtnt2d5tbTWQWfmldbXnayUQmjFt_NUuNbQY2LabY' \
  -H 'x-user-id: 1'