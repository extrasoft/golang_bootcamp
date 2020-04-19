package main

import (
	"fmt"
	"myFmt"
	"log"
	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println("Hello World")
	myFmt.Println("Test")
	log.Println("Test Log")
	stringutil.Reverse("Helloworld")
}
