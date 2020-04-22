package main

import "fmt"

func main() {
	x := 5

	if y := x + 1; (x <= 5) && false {
		fmt.Println("if :", y)
	} else if x <= 6 {
		fmt.Println("else if :", y)
	}else{
		fmt.Println("else :", y)
	}

	// fmt.Println(y) // error เพราะตัวแปร y อยู่ด้านนอกของ if else
}
