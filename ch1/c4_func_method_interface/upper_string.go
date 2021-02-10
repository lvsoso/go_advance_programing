package main

import (
	"os"
	"strings"
	"fmt"
)

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

//type fmt.Stringer interface {
//String() string
//}

func main() {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}