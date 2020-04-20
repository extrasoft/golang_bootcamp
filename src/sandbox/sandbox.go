package main

import (
	"fmt"
)

type KG float64
type LB float64

func main() {
	x := KG(3)
	y := LB(3)

	fmt.Println(x == y)     // error เพราะถือเป็นคนละ type กัน
	fmt.Println(x == KG(y)) // เปรียบเทียบได้เพราะ เราทำการ cast type ของ y ให้เป็น KG
	fmt.Println(x == 3)   // เปรียบเทียบได้เพราะ 3 ถือเป็นค่า constant และเป็นค่าที่สามารถแปลงเป็น float64 ได้
	

	result := x + 5
	fmt.Println("result:", result) // เปรียบเทียบได้ เพราะ 5 เป็นค่าที่สามารถแปลงเป็น float64 ได้

	result2 := x + "5" // error เพราะ "5" เป็นค่าที่ไม่สามารถแปลงเป็น float64 ได้
	fmt.Println("result2:", result2)
}