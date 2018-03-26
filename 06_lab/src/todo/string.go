package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randCharSet = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randStringRunes(n int) string {
	b := make([]rune, n)
	chars := len(randCharSet)
	for i := range b {
		b[i] = randCharSet[rand.Intn(chars)]
	}
	return string(b)
}
