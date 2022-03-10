package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var open []int

	if len(os.Args) > 1 {
		ip := os.Args[1]
		for i := 1; i <= 66534; i++ {
			IsOpen := scan(ip, i)

			if IsOpen {
				open = append(open, i)
				fmt.Printf("Open ports: %d \n", open)
			}
		}
	} else {
		log.Fatal("please provide ip address")
	}

}
