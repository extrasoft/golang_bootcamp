package main

import (
	"fmt"
)

func main() {
	
	x := 70
	var p *int	 // ประกาศตัวแปร Pointer (*) ชนิด int มีค่า default = nil
	p = &x		 // กำหนดให้ value ของ p = address ของ x

	fmt.Println("Value x : ", x)
	fmt.Println("Address &x : ", &x)		// &x จะเป็นการอ้างอิงถึง Address ของตัวแปร x
	fmt.Println("dereference Value *(&x)  : ", *(&x)) // *(&x) จะเป็นการ dereference ค่าของตัวแปร x (x ไม่ใช่ตัวแปร pointer)
	fmt.Println("---------------------")
	fmt.Println("Value p : ", p)
	fmt.Println("Address &p : ", &p)
	fmt.Println("dereference Value p  : ", *p)	// *p จะเป็นการ dereference ค่าของตัวแปร x ที่ตัวแปร p เก็บไว้
}
