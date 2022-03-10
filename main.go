package main

import (
	"fmt"
	"os"
)

func main() {
	ip := os.Args[1]
	var open []int

	// looping all ports
	for i := 1; i <= 66534; i++ {
		IsOpen := scan(ip, i)

		if IsOpen {
			open = append(open, i)
			fmt.Printf("Open ports: %d \n", open)
		}
	}

}
