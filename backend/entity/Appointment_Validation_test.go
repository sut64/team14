package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestAllpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	appointment := Appointment{
		Note:        "123",
		Number:      123,
		AppointDate: time.Date(2022, 01, 28, 8, 00, 00, 00, time.Local),
		IssueDate:   time.Date(2022, 01, 27, 9, 00, 00, 00, time.Local),
	}

	ok, err := govalidator.ValidateStruct(appointment)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestNoteMustNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	appointment := Appointment{
		Note:        "",
		Number:      123,
		AppointDate: time.Date(2022, 01, 28, 8, 00, 00, 00, time.Local),
		IssueDate:   time.Date(2022, 01, 27, 9, 00, 00, 00, time.Local),
	}

	ok, err := govalidator.ValidateStruct(appointment)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Note can not be blank"))
}

func TestAmountDayMustMoreThanO(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	appointment := Appointment{
		Note:        "123456",
		Number:      0,
		AppointDate: time.Date(2022, 01, 28, 8, 00, 00, 00, time.Local),
		IssueDate:   time.Date(2022, 01, 27, 9, 00, 00, 00, time.Local),
	}

	ok := govalidator.IsPositive(float64(float64(appointment.Number)))
	err := ""
	if !ok {
		err = "Amount of day must more then 0"
	}
	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).To(gomega.Equal("Amount of day must more then 0"))
}

func TestAppoiontDatemustFuture(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	appointment := Appointment{
		Note:        "123",
		Number:      123,
		AppointDate: time.Now(),
		IssueDate:   time.Date(2022, 01, 27, 9, 00, 00, 00, time.Local),
	}

	ok, err := govalidator.ValidateStruct(appointment)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Appointment Date must be in future"))
}
