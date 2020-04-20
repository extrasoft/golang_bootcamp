package main

import "fmt"

func fp() *int {
	x := 4
	return &x
}

func main() {
	a := fp()
	fmt.Printf("Type of a = %T\n", a)
	fmt.Println("a = ", a)
	fmt.Println("*a = ", *a)

	fmt.Println("a == fp(): ", a == fp()) 
	// ผลลัพธ์ false เพราะว่าทุกครั้งที่เรียก fp() จะเป็นการสร้าง local variable ตัวใหม่ address ใหม่
	
	fmt.Println(a, fp())
}