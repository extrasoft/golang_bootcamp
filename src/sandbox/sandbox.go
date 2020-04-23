package main

import (
	"fmt"
	"math/rand"
	"prime"
)

func main() {
	for i := 1; i <= 30; i++ {
		rn := rand.Intn(1000000)
		fmt.Printf("%d %t\n", rn, prime.IsPrime(rn))
	}
}
