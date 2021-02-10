package main

import (
	"fmt"
	"unicode/utf8"
)

func str2runes(s string) []rune{
	var p []int32
	for len(s)>0 {
		r,size:=utf8.DecodeRuneInString(s)
		p=append(p,int32(r))
		s=s[size:]
	}
	return []rune(p)
}

func main() {
	fmt.Println(str2runes("世界abc"))
}
