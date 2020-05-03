package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	// เปลี่ยนรูปแบบของ log
	log.SetFlags(log.Ltime)

	// สร้าง server และเปิด port 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		// ถ้ามี error ก็จะพ่น log แล้วจบการทำงานเลยเพราะไปต่อไม่ได้แล้ว
		log.Fatal(err)
	}

	// สร้าง for infinite loop เพื่อที่จะสามารถรับ connection ใหม่ได้เรื่อยๆ
	for {
		// สร้าง connection
		conn, err := listener.Accept()

		// log เพื่อบอกว่ามีคนเข้ามา connect ใหม่
		// conn.RemoteAddr() คือระบุ ip และ port ของ client
		log.Println("New client is connected : ", conn.RemoteAddr())

		if err != nil {
			// ถ้า connection ถูก close โดย client ก็จะพ้น log ออกมา
			log.Println(err)
			return
		}

		// ใช้ goroutine เพือให้แต่ละ client ไม่ต้องรอซึ่งกันและกัน (มี go routine ของแต่ลคน)
		go countingDownHandler(conn)
	}
}

/*
 	telnet localhost 8080
 	ถ้าต้องการหยุด telnet ต้องพิมพ์ ctrl+]
	จากนั้นกด q
	type เป็น interface
*/
func countingDownHandler(conn net.Conn) {

	// ทำเป็นอย่างสุดท้าย
	defer func() {
		// มาถึงบรรทัดนี้ก่อนปิด connection จะเขียน string มาก่อนชุดนึง
		io.WriteString(conn, "Your connection will be closed by server\n")
		conn.Close()
	}() // annonymous function
	io.WriteString(conn, "Enter number : ")

	// รับ input จาก client
	input := bufio.NewScanner(conn)
	count := Scan(input)

	for {
		io.WriteString(conn, strconv.Itoa(count)+"\n")

		time.Sleep(time.Second) // หน่วงเวลา 1 วินาทีใน loop แต่ละรอบ
		count--

		// ทดสอบการ close connection by server
		// เมื่อ client ใส่ค่ามากกว่า 10
		if count > 10 {
			return
		}

		if count < 0 {
			io.WriteString(conn, "Enter number : ")
			count = Scan(input)
			if count == 0 {
				break
			}
		}
	}

}

func Scan(input *bufio.Scanner) int {

	if ok := input.Scan(); !ok {
		log.Println("Cannot scan value from client")
		log.Println("Connection is closed by client")
		// return กลับไปแต่ยังไม่หยุดการทำงานของ program
		return 0
	}

	// รับ input จาก client ที่เป็น text มาแปลงเป็น int
	count, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("Cannot convert value from Text to int", input.Text())
		return 0
	}
	return count
}
