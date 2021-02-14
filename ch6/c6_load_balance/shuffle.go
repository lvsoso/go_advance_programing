package main

import "math/rand"

func shuffle(n int) []int {
	b := rand.Perm(n)
	return b
}
