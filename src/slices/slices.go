package main

import "fmt"

func main() {
	// แบบที่สาม make(ชนิดของอาเรย์, len, cap)
	a := make([]int, 10, 11)
	fmt.Println(len(a), cap(a), a) 
	// 10 11 [0 0 0 0 0 0 0 0 0 0]

	a[9] = 5
	fmt.Println(len(a), cap(a), a) 
	// 10 11 [0 0 0 0 0 0 0 0 0 5]

	a = append(a, 10)
	fmt.Println(len(a), cap(a), a) 
	// 11 11 [0 0 0 0 0 0 0 0 0 5 10]

	a = append(a, 12)
	fmt.Println(len(a), cap(a), a) 
	// 12 22 [0 0 0 0 0 0 0 0 0 5 10 12]
}

/*	วิธีที่ 1
func main() {
	fruits := [...]string{"apple", "banana", "papaya", "orange", "mango", "pineapple"}
	fmt.Println("fruits", len(fruits), cap(fruits), fruits)

	// หั่นตั้งตำแหน่งที่ 1 แต่ไม่เกิน 4
	myFavor := fruits[1:4]
	fmt.Println("before myFavor", len(myFavor), cap(myFavor), myFavor)

	// copy slices
	yourFavor := myFavor

	// การเปลี่ยนแปลงค่าของ slice ตัวใหม่จะส่งผลต่อ slice ตัวเก่าด้วยเพราะว่าทั้งสอง slice เก็บอยู่ในที่เดียวกัน
	yourFavor[0] = "guava"
	fmt.Println("after myFavor", len(myFavor), cap(myFavor), myFavor)
	fmt.Println("copy slice", len(yourFavor), cap(yourFavor), yourFavor)
}
*/