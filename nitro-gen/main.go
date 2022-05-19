package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandStringRunes(n int, letterRunes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Code() string {
	var code string
	var dict = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := 0; i < 16; i++ {
		code = code + RandStringRunes(1, dict)
	}
	return code
}

func main() {

	n := 0
	start := time.Now()
	for n < 99999 {
		fmt.Printf("discord.com/gift/%s\n", Code())
		n += 1
	}
	elapsed := time.Since(start)
	fmt.Printf("  [ > ] TTC: %s\n\n",  elapsed)
}
