package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestPAllpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	prevention := Prevention{
		Disease:    "กก",
		Protection: "กข",
		Age:        10,
		Date:       time.Now(),
	}

	ok, err := govalidator.ValidateStruct(prevention)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestDiseaseMustNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	prevention := Prevention{
		Disease:    "",
		Protection: "กข",
		Age:        10,
		Date:       time.Now(),
	}

	ok, err := govalidator.ValidateStruct(prevention)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Disease cannot be blank"))
}

func TestProtectionMore5(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	prevention := Prevention{
		Disease:    "กก",
		Protection: "ก",
		Age:        10,
		Date:       time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(prevention)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Protection must be more than 5"))
}

func TestAgeMustPositive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	prevention := Prevention{
		Disease:    "กก",
		Protection: "กข",
		Age:        -10,
		Date:       time.Now(),
	}

	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(prevention)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Age does not validate as positive"))
}

func TestPreventionDateNotBeFuture(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	prevention := Prevention{
		Disease:    "กก",
		Protection: "กข",
		Age:        10,
		Date:       time.Now().Add(time.Hour * 24), // ผิด เป็นอนาคต
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(prevention)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Date cannot be in the future"))
}
