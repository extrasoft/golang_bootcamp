package main

import (
	"fmt"
	"reflect"
)

// parameter ใน function ของ golang จะไม่สามารถใส่ค่า default value ได้
func avg(a, b float64) (c float64) {
	c = (a + b) / 2
	return c
}


func main() {
	a := avg(1, 3)
	b := avg(1, 2)
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(avg))	// ไว้สำหรับดู parameter และ return ของ function
}