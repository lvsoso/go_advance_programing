```shell
sudo /usr/local/go/bin/go install -buildmode=shared std
sudo /usr/local/go/bin/go build -linkshared -o hello_linkshared hello.go

```