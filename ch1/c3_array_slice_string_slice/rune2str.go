package main

import (
	"fmt"
	"unicode/utf8"
)

func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}

func main() {
	fmt.Println(runes2string([]int32{19990, 30028, 97, 98, 99}))
}
