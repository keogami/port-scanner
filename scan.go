package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//function to scan the ip and port
func scan(ip string, port int) bool {
	return rand.Intn(2) == 0
}

//main function

func main() {
	// looping all ports
	result := 0
	for i := 1; i <= 66534; i++ {
		result = scan("127.0.0.1", i)
		// result = i
		result += 1
		fmt.Println(result)
	}
}
