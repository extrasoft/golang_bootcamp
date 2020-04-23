package prime

import (
	"fmt"
	"math"
)

var notPrimes [1000000]bool

func init() {
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrimes)))))
	fmt.Println("sqrtLen = ", sqrtLen)
	for i := 2; i < sqrtLen; i++ {
		if notPrimes[i] {
			continue
		}
		notPrimes[i] = false
		for j := i * 2; j < len(notPrimes); j += i {
			notPrimes[j] = true
		}
	}
	fmt.Println("Initialized")
}

func IsPrime(i int) bool {
	return !notPrimes[i]
}
