package main

import (
	"fmt"
)

func main() {
	// กำหนดค่า default ให้กับ const แบบ group
	const (
		Sunday = 1
		Monday
		Tuesday
		Wednesday = 2
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
