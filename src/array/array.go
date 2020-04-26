package main

import "fmt"

func main() {
	// 	แบบที่ 1
	fruitsOne := [5]string{}
	fmt.Println(len(fruitsOne), fruitsOne)

	// แบบที่ 2
	fruitsTwo := [5]string{
		"apple",
		"banana",
		"orange",
		"mango",
	}
	fmt.Println(len(fruitsTwo), fruitsTwo)

	// แบบที่ 3 แบบ elixir ใช้ในกรณีไม่ทราบจำนวน element ที่แน่นอน
	fruitsThree := [...]string{
		"apple",
		"banana",
		"orange",
		"mango",
	}
	fmt.Println(len(fruitsThree), fruitsThree)

	// แบบที่ 4 แบบ กำหนด index
	fruitsFour := [...]string{
		3: "apple",
		1: "banana",
		2: "orange",
		0: "mango",
	}
	fmt.Println(len(fruitsFour), fruitsFour)

	// แบบที่ 5 แบบ กำหนด index
	fruitsFive := [...]string{
		"apple",
		20: "banana",
		"orange",
		"mango",
	}
	fmt.Println(len(fruitsFive), fruitsFive)

	// แบบที่ 6 แบบ array 2 มิติ
	fruitsSix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(len(fruitsSix), fruitsSix)

	fruitsSeven := [...] int {}
	fmt.Println(len(fruitsSeven), fruitsSeven)
}
