package weight

type KG float64
type LB float64

/* 
	การประกาศตัวแปร
		"ขึ้นต้นเป็นตัวพิมพ์เล็ก" จะสามารถเรียกใช้ภายใน Package เดียวกันเท่านั้น ไม่สามารถเรียกใช้ข้าม Package ได้ (Private)
		"ขึ้นต้นเป็นตัวพิมพ์ใหญ่" จะสามารถเรียกใช้ข้าม Package ได้ (Public)
*/
const kgToLBRatio = 2.20462262
