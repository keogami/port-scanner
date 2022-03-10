package main

import "fmt"

func main() {
	var open []int

	// looping all ports
	for i := 1; i <= 66534; i++ {
		result := scan("127.0.0.1", i)

		if result != false {
			openPorts := append(open, i)
			fmt.Printf("Open ports: %d \n", openPorts)
			// for j := range openPorts {
			// 	fmt.Printf()
			// }
		}
	}

}
