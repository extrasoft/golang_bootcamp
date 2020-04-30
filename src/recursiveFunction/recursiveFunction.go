package main

import "fmt"

func countDown(c int) {
	fmt.Println(c)
	if c == 0 {
		// ออกจาก functio
		return
	}
	countDown(c - 1)
}

func main() {
	countDown(5)
}
