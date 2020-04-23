package main

import "fmt"

func main() {
	// for init; condition; increment {}
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // next loop
		}
		fmt.Println(i)
		if i == 4 {
			break // exit for loop
		}
	}

}
