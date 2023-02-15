package entity

import (
	// . "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~ห้ามว่าง"` // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string `valid:"matches(^[LMH][0-9]{7}$)"` // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}

// func TestOne(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	c := Customer{}
// }
