package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ทอสอบข้อมูลถูกต้อง
func TestContagiousPass(t *testing.T) {
	g := NewGomegaWithT(t)

	contagious := Contagious{
		Name:       "Abcde",
		Symptom:    "asdfgjkll vjkklp",
		Incubation: 7,
		Date:       time.Now(),
	}

	ok, err := govalidator.ValidateStruct(contagious)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())
}

//ทดสอบชื่อเป็นค่าว่าง ต้องเจอ error
func TestContagiousNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	contagious := Contagious{
		Name:       "", //ผิด
		Symptom:    "asdfgjkll vjkklp",
		Incubation: 7,
		Date:       time.Now(),
	}

	ok, err := govalidator.ValidateStruct(contagious)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name cannot be blank"))

}

