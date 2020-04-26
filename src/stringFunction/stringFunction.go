package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode"
)

func main() {
	// strconv
	atoi, _ := strconv.Atoi("999") 	// แปลง ตัวอักษร => ตัวเลข
	itoa := strconv.Itoa(123)		// แปลง ตัวเลข => ตัวอักษร
	fmt.Println(atoi, reflect.TypeOf(atoi))
	fmt.Println(itoa, reflect.TypeOf(itoa))

	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.ParseBool("fAlse"))
	fmt.Println(strconv.ParseBool("0"))

	// unicode
	fmt.Println("unicode.IsDigit('1')", unicode.IsDigit('1'))
	fmt.Println("unicode.IsDigit('๑')", unicode.IsDigit('๑'))
	fmt.Println("unicode.IsNumber('5')", unicode.IsNumber('5'))
	fmt.Println("unicode.IsUpper('a')", unicode.IsUpper('a'))
	fmt.Println("unicode.IsUpper('A')", unicode.IsUpper('A'))
	fmt.Println("unicode.In('ข', unicode.Thai)", unicode.In('ข', unicode.Thai))
}
