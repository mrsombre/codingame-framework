package main

import (
	"math/rand"
	"runtime"
	"time"
)

var rnd *rand.Rand

func init() {
	runtime.GOMAXPROCS(1)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	debug = true
}
