package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type UpperWriter struct {
	io.Writer
}

func (w *UpperWriter) Write(data []byte) (n int, err error) {
	return w.Writer.Write(bytes.ToUpper(data))
}

func main(){
	fmt.Fprint(&UpperWriter{os.Stdout}, "hello, world")
}