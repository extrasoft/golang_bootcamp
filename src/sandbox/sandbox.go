package main

import "fmt"

var z = 0
func main() {
	// แบบที่ 1 : var name type = expression
	var a int = 1 + 2
	fmt.Println("a = ", a)

	// แบบที่ 2 : var name = expression
	var b = 4
	var c = 4.0
	var d = "4"
	fmt.Printf("b = %#v = %T\n", b, b) // %T หมายถึง แสด Type ของตัวแปร
	fmt.Printf("c = %#v = %T\n", c, c) // %#v หมายถึง แสดง value ของตัวแปร
	fmt.Printf("d = %#v = %T\n", d, d)

	// แบบที่ 3 : var name type
	var e int
	var f float32
	var g string
	var h []string // เรียกว่า slice
	fmt.Printf("e = %#v = %T\n", e, e)
	fmt.Printf("f = %#v = %T\n", f, f)
	fmt.Printf("g = %#v = %T\n", g, g)
	fmt.Printf("h = %#v = %T\n", h, h)
	fmt.Println("h == nil : ", h == nil) // true

	// แบบที่ 4 (Short Declaration) ไม่สามารถประกาศนอก func ได้: name := expression
	i := 10 + 5
	j, k := 2, 3
	fmt.Printf("i = %#v = %T\n", i, i)
	fmt.Printf("j = %#v = %T\n", j, j)
	fmt.Printf("k = %#v = %T\n", k, k)

	/*** ประกาศตัวแปรแล้วประกาศซ้ำไม่ได้ ***/
	fmt.Println("z = ", z)
}