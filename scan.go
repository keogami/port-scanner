package main

import (
	"math/rand"
	"time"
)

func init() {
  rand.Seed(time.Now().Unix())
}

func scan(ip string, port int) bool {
  return rand.Intn(2) == 0
}
