package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~ห้ามว่าง"` // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string `valid:"matches(^[LMH][0-9]{7}$)~กรอกรูปแบบให้ถูกต้อง,required~ห้ามว่าง"` // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}

func TestTruecase(t *testing.T) {
	g := NewGomegaWithT(t)

	c := Customer{
		Name:       "Taweechai",
		Email:      "tawee@gmail.com",
		CustomerID: "M0252526",
	}

	ok, err := govalidator.ValidateStruct(c)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	//ถูกทั้งหมด ไม่จำเป็นต้องมี error
}
