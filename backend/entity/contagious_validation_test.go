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
// ทดสอบอาการน้อยกว่า 10 ตัวอักษร ต้องเจอ error
func TestContagiousSymptomMore10(t *testing.T) {
	g := NewGomegaWithT(t)

	contagious := Contagious{
		Name:       "Abc",
		Symptom:    "ab", // ผิด
		Incubation: 7,
		Date:       time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(contagious)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Symptom must be more than 10"))
}

// ทดสอบระยะฟักตัวเป็นเลขที่ไม่ได้อยู่ระหว่าง 1 - 90 ต้องเจอ error
func TestContagiousIncubationNotMinus(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []int{
		-7,
		99,
	}
	for _, fixture := range fixtures {
		contagious := Contagious{
			Name:       "Abc",
			Symptom:    "abcdefg hijk",
			Incubation: fixture, // ผิด
			Date:       time.Now(),
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(contagious)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Incubation must be between 1-90"))
	}
}

