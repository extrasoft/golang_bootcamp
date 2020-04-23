package main

import (
	"fmt"
	"weight"
)

func main() {
	w := weight.KG(1)
	fmt.Println("KG:", w)
	fmt.Println("KG to LB:", weight.KgToLB(w))
}
