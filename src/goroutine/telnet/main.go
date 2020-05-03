package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Ltime)
	// สร้าง custom telnet ขึ้นมาเพื่อที่จะสามาร custom การแสดงผลได้
	// วิธีใช้ต้อง เปิด server ที่ go run goroutine\countdownserver ก่อน
	// จากนั้นใช้ go run goroutine\telnet
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	log.Println("Connected to server successfully.")

	// copy ข้อความจาก connection ของ server มาที่ client
	// ต้องทำเป็น go routine เพื่อที่จะไม่ต้องรอ copy ข้อความจาก server มา client
	// log.Println("Copy from conn to stdout")
	go io.Copy(os.Stdout, conn)

	// copy ข้อความจาก connection ของ client ไปที่ server
	// log.Println("Copy from stdin to conn")
	io.Copy(conn, os.Stdin)
}
