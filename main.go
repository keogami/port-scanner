package main

import "fmt"

func main() {
	// looping all ports
	// result := 0
	for i := 1; i <= 66534; i++ {
		result := scan("127.0.0.1", i)
		// result = i
		// result += 1
		fmt.Printf("Port number %d open? => %t \n", i, result)
	}
}
