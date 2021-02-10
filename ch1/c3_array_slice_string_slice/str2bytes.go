package main

import "fmt"

func str2bytes(s string)[]byte {
	p := make([]byte, len(s))
	for i:= 0; i < len(s); i ++ {
		c := s[i]
		p[i] = c
	}
	return p
}

func main() {
	fmt.Println(str2bytes("世界abc"))
}
