package main

import (
	gla2 "golangAss2/pckg"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	gla2.GenerateRandomTxs(10)
}