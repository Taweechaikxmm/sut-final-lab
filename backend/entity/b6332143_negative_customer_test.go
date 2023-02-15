package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	c := Customer{
		Name:       "Taweechai",
		Email:      "tawee@gmail.com",
		CustomerID: "M025252", //ผิด
	}

	ok, err := govalidator.ValidateStruct(c)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("กรอกรูปแบบให้ถูกต้อง"))
}

func TestNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	c := Customer{
		Name:       "Taweechai",
		Email:      "tawee@gmail.com",
		CustomerID: "", //ผิด
	}

	ok, err := govalidator.ValidateStruct(c)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("ห้ามว่าง"))
}
