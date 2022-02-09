package entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestMVAllpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	medicineandvaccine := MedicineandVaccine{
		RegNo:  "A1234",
		Name:   "Pizer",
		MinAge: 18,
		MaxAge: 90,
		Date:   time.Now().Add(24 * time.Hour),
	}
	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(medicineandvaccine)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

}
func TestRegNoMustBeInValidPattern(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	fixtures := []string{
		"a000",
		"b0b0",
		"c000",
		"000a",
		"a00000",
	}
	for _, fixtures := range fixtures {
		medicineandvaccine := MedicineandVaccine{
			RegNo:  "a123",
			Name:   "Pizer",
			MinAge: 18,
			MaxAge: 90,
			Date:   time.Now(),
		}

		//ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(medicineandvaccine)

		//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(gomega.BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(gomega.BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(gomega.Equal(fmt.Sprintf(`RegNo %s does not validate as  matches(^[A-Z]{1}\\d{4}$)`, fixtures)))
	}
}
func TestNameMustNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	medicineandvaccine := MedicineandVaccine{
		RegNo:  "A1234",
		Name:   "",
		MinAge: 18,
		MaxAge: 90,
		Date:   time.Now().Add(24 * time.Hour),
	}
	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(medicineandvaccine)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank"))
}

func TestMinAgeMustPositive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	medicineandvaccine := MedicineandVaccine{
		RegNo:  "A1234",
		Name:   "Pizer",
		MinAge: -18,
		MaxAge: 90,
		Date:   time.Now(),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(medicineandvaccine)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("MinAge does not validate as positive"))
}

func TestMaxAgeMustPositive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	medicineandvaccine := MedicineandVaccine{
		RegNo:  "A1234",
		Name:   "Pizer",
		MinAge: 18,
		MaxAge: -90,
		Date:   time.Now(),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(medicineandvaccine)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("MaxAge does not validate as positive"))
}

func TestDateMustNotPast(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	medicineandvaccine := MedicineandVaccine{
		RegNo:  "A1234",
		Name:   "Pizer",
		MinAge: 18,
		MaxAge: 90,
		Date:   time.Now().Add(20 - time.Hour),
	}

	ok, err := govalidator.ValidateStruct(medicineandvaccine)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Date must not be past"))
}

