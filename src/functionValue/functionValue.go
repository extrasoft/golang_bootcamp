package main

import (
	"fmt"
	"math"
)

func add(a float64, b float64) float64 {
	return a+b
}

func sub(a float64, b float64) float64 {
	return a-b
}

func apply(a,b float64, fn func(float64, float64) float64) (float64, error) {
	// ถ้าไม่ส่งชื่อ function มาจะ return error propagate กลับไป
	if fn == nil {
		// math.NaN() จะ return ค่า NaN
		return math.NaN(), fmt.Errorf("apply: nil operation")
	}

	return fn(a, b), nil
}

func main() {
	a, _ := apply(1,2,add)
	b, _ := apply(1,2,sub)
	c, _ := apply(1,2,nil)
	fmt.Println(a, b, c)
}