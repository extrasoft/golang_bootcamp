package main

import "fmt"

func main() {
	x := "f"
	// switch {....} => switch true {.....}
	switch x {
	case "a":
		fmt.Println("a")
		fallthrough
	case "b":
		fmt.Println("b")
	case "c", "d", "e", "f":
		fmt.Println("c", "d", "e", "f")
		if x == "f" {
			break;
		}
		fmt.Println("You are lucky")
	case "z":
	default:
		fmt.Println("Case Default")
	}

}
