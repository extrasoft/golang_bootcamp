package main

import "fmt"

// global variable
// var x = 0

// return เป็น function ที่ไม่รับอะไรเข้ามาเลย
// และ function นั้นก็ return int อีกทีนึง
func createF() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func main() {
	// วิธีเรียกใช้ annonymous function จะคล้ายๆกับ func()() 
	// โดย () แระจะหมายถึงเรียกใช้ function
	// () ตัวที่สองจะหมายถึง execute function ภายในตัวมันอีกทีนึง
	// เพราะว่าเราทำการสร้าง function ที่ return เป็น function

	f := createF()		// เรียกใช้ function ครั้งแรก
	fmt.Println(f())	// execute function => 1
	fmt.Println(f())	// execute function => 2
	fmt.Println(f())	// execute function => 3

	f2 := createF()		// เรียกใช้ function ครั้งสอง
	fmt.Println(f2())	// execute function => 1
	fmt.Println(f2())	// execute function => 2
	fmt.Println(f2())	// execute function => 3
}