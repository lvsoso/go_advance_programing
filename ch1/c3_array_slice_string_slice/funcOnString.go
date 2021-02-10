package main

import (
	"fmt"
	"unicode/utf8"
)

func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}

func forBody(i int, r rune)  {
	fmt.Println(i, r)
}

func main() {
	forOnString("世界abc",forBody)
}
