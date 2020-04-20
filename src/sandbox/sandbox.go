package main

import "fmt"

func multiple(y *int) {
	// รับค่าเข้ามาเป็น pointer
	// ทำการ dereference โดยใช้ *ตัวแปร
	*y = *y * 2
}

func main() {
	x := 2
	fmt.Println("x before :",x)
	multiple(&x)	//ส่งเป็น address ของ x
	fmt.Println("x after :",x)
}