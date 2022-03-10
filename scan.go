package main

import (
  "math/rand"
)

func scan(ip string, port int) bool {
  return rand.Intn(2) == 0
}
