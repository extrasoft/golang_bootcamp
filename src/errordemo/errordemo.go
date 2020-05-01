package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const dbReady = false
const balance = 200
const numberToSuccess = 100000 // จำนวนที่จำลอง connect สำเร็จ

func connectDB(nTry int) error {
	if nTry == numberToSuccess {
		return nil
	}

	return errors.New("busy")
}

func waitForDatabase() error {
	timeout := 3 * time.Second
	deadline := time.Now().Add(timeout) // กำหนดเวลาสิ้นสุดในการ connect เป็น 3 วินาที

	// วนลูปเพื่อกำหนดเวลาในการลอง connect ก่อนหมดเวลา
	// เช่น มีเวลาในการลอง connect กี่ครั้งก็ได้ภายใน 3 วิ
	for tries := 0; time.Now().Before(deadline); tries++ {
		err := connectDB(tries)
		if err == nil {
			return nil
		}
		log.Printf("%v database is not responding (%v); retrying...",tries, err)
		
		// ถ้าต้องการให้การ trying แต่ละครั้งหน่วงเวลา 1 วินาที
		// time.Sleep(time.Second)
		// หรือ
		// time.Sleep(time.Second << uint(tries))	// 1 2 4 8
	}

	// ถ้าหมดเวลาแล้วยัง connect ไม่ได้ก็จะปริ้น log แล้วออกจากการทำงาน
	return fmt.Errorf("waitForDatabase: timeout %v", timeout)
}

func getBalance() (int, error) {
	if !dbReady {
		// log.Println("retrying....")
		if err := waitForDatabase(); err != nil {
			return 0, fmt.Errorf("getBalance(), %v", err)
		}
		// return 0, errors.New("withdraw: database is down")
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
	log.SetFlags(0) // 0 คือไม่มี pattern log เลย
	amount, err := withdraw(200)
	if err != nil {

		// fmt.Println("Main(): ", err)
		// os.Exit(1)	//	os.Exit(1) หมายถึงออกจากโปรแกรม

		// log.Fatalf มีค่าเท่ากับ 2 บรรทัดบนคือ print log และ ออกจากโปรแกรมเลย
		log.Fatalf("Main(): %v", err)
	}
	fmt.Println("Please collect your money", amount)
}
