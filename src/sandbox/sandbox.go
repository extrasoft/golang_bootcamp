package main

import (
	"fmt"
	// "weight"
	"github.com/extrasoft/golang_bootcamp/src/weight"
)

func main() {
	w := weight.KG(1)
	fmt.Println("KG:", w)
	fmt.Println("KG to LB:", weight.KgToLB(w))
}
