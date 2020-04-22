package main

import (
	"fmt"
)

type KG float64
type LB float64

func (lb LB) toKG() KG {
	return KG(lb / 2.2046226218)
}

func (kg KG) toLB() LB {
	return LB(kg * 2.2046226218)
}

func main() {
	x := KG(1)
	y := LB(2.2046226218)

	fmt.Println(x == y.toKG())
	fmt.Println(y == x.toLB())

}
