package main

import "fmt"

func main() {
	testValue("123")
	testValue(123)
	testValue(true)
	testValue(123.45)
	testValue(float32(123.45))
}

func testValue(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("string value: ", v)
	case int:
		fmt.Println("int value: ", v)
	case float32, float64:
		fmt.Println("float value: ", v)
	case bool:
		fmt.Println("boolean value: ", v)
	default:
		fmt.Println("No match any")
	}
}
