package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//function to scan the ip and port
func scan(ip string, port int) bool {
	_, _ = ip, port
	return rand.Intn(2) == 0
}
