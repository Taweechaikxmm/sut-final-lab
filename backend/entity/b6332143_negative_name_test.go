package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameNotBeBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	c := Customer{
		Name:       "", //ผิด
		Email:      "tawee@gmail.com",
		CustomerID: "M0252526",
	}

	ok, err := govalidator.ValidateStruct(c)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("ห้ามว่าง"))
}
