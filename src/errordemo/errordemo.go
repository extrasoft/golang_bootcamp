package main

import (
	"errors"
	"fmt"
	"log"
)

const dbReady = true
const balance = 200

func getBalance() (int, error) {
	if !dbReady {
		return 0, errors.New("withdraw: database is down")
	}
	return balance, nil
}



func withdraw(amount int) (int, error) {
	balance, err := getBalance()
	if err != nil {
		// fmt.Errorf จะคล้ายๆกับ printf คือสามารถส่งค่า parameter เข้าไปใน function ได้สองตัว (สามารถใส่ตัวแปรไปได้)
		// แต่ errors.New() จะใส่ได้แค่ตัวเดียว
		return 0, fmt.Errorf("withdraw %v", err)
	}
	if amount > balance {
		return 0, errors.New("withdraw: Insuficient fund")
	}
	return 100, nil
}

func main() {
	log.SetFlags(0)	// 0 คือไม่มี pattern log เลย
	amount, err := withdraw(200)
	if err != nil {
		
		// fmt.Println("Main(): ", err)
		// os.Exit(1)	//	os.Exit(1) หมายถึงออกจากโปรแกรม

		// log.Fatalf มีค่าเท่ากับ 2 บรรทัดบนคือ print log และ ออกจากโปรแกรมเลย
		log.Fatalf("Main(): %v", err)
	}
	fmt.Println("Please collect your money", amount)
}
