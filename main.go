package main

import "fmt"

func main() {
	var open []int

	// looping all ports
	for i := 1; i <= 66534; i++ {
		IsOpen := scan("127.0.0.1", i)

		if IsOpen {
			open = append(open, i)
			fmt.Printf("Open ports: %d \n", open)
		}
	}

}
