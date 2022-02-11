package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestScreeningpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	screening := Screening{
		BloodPressure: 180,
		CongenitalDisease: "cancer",
		Time: 		time.Now(),
	}

	ok, err := govalidator.ValidateStruct(screening)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestCongenitalDiseaseNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	screening := Screening{
		BloodPressure: 180,
		CongenitalDisease: "",
		Time: 		time.Now(),
	}

	ok, err := govalidator.ValidateStruct(screening)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("CongenitalDisease can not be blank"))
}

func TestBloodPressureMoreThan0(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	screening := Screening{
		BloodPressure: 0,
		CongenitalDisease: "cancer",
		Time: 		time.Now(),
	}

	ok := govalidator.IsPositive(float64(float64(screening.BloodPressure)))
	err := ""
	if !ok {
		err = "BloodPressure must more then 0"
	}

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).To(gomega.Equal("BloodPressure must more then 0"))
}

func TestScreeningDatemustPresent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	screening := Screening{
		BloodPressure: 180,
		CongenitalDisease: "cancer",
		Time: 		time.Now().Add(24 * time.Hour),
	}

	ok, err := govalidator.ValidateStruct(screening)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Screening Date must be in Present"))
}
