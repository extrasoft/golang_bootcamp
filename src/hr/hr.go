package hr

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person	// anonymous field คือ ไม่ระบุชื่อ field แต่ว่าใส่ struct ลงไปเลย
	Description string
}