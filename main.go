package main

import "fmt"

func main() {
	// looping all ports
	for i := 1; i <= 66534; i++ {
		result := scan("127.0.0.1", i)
		fmt.Printf("Port number %d open? => %t \n", i, result)
	}
}
