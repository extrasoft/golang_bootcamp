package main

import "fmt"

func main() {
	// ประกาศแบบวิธีแรก map[keyType]valueType
	items := map[string]int{
		"pen":    4,
		"pencil": 5,
		"rubber": 3,
		"a":      1,
		"b":      2,
	}
	fmt.Println(items)           // map[a:1 b:2 pen:4 pencil:5 rubber:3]
	fmt.Println(items["pencil"]) // 5

	// ประกาศแบบวิธีที่สอง make(map[keyType]valueType)
	orders := make(map[string]int)
	orders["ruler"] = 6

	fmt.Println("Before orders: ", orders)          // Before orders:  map[ruler:6]
	fmt.Println("Before orders: ", orders["ruler"]) // Before orders:  6

	// copy maps จะส่งผลให้ตัวที่ copy มาเปลี่ยนค่าด้วยเพราะตัวแปรเก็บไว้ตำแหน่งเดียวกัน
	newOrders := orders
	newOrders["ruler"] = 8
	fmt.Println("After orders: ", orders)       // After orders:  map[ruler:8]
	fmt.Println("After newOrders: ", newOrders) // After newOrders:  map[ruler:8]

	// วิธีตรวจสอบว่ามี key อยู่ใน maps หรือไม่
	if v, ok := items["xxx"]; !ok {
		fmt.Println("xxx not exist") // xxx not exist
	} else {
		fmt.Println("xxx exist : ", v)
	}

	// วิธีวนลูปใน maps โดยใช้ for....range
	for k, v := range items {
		fmt.Println(k, v)
	}
	// pen 4
	// pencil 5
	// rubber 3
	// a 1
	// b 2
}
