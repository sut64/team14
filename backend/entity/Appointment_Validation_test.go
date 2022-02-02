package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

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
