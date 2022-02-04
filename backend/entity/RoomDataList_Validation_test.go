package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	roomdatalist := RoomDataList{
		Note:          "DONTDISTURB",
		Day:           9,
		EnterRoomTime: time.Date(2022, 01, 27, 9, 00, 00, 00, time.Local),
	}

	ok, err := govalidator.ValidateStruct(roomdatalist)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestNoteNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	roomdatalist := RoomDataList{
		Note:          "",
		Day:           999,
		EnterRoomTime: time.Date(2022, 01, 29, 9, 00, 00, 00, time.Local),
	}

	ok, err := govalidator.ValidateStruct(roomdatalist)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Note can not be blank"))
}

func TestAmountDayMustMoreThanO(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	roomdatalist := RoomDataList{
		Note:          "999999999",
		Day:           0,
		EnterRoomTime: time.Date(2022, 01, 29, 9, 00, 00, 00, time.Local),
	}

	ok := govalidator.IsPositive(float64(float64(roomdatalist.Day)))
	err := ""
	if !ok {
		err = "Amount of day must more then 0"
	}
	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).To(gomega.Equal("Amount of day must more then 0"))
}

func TestEnterRoomTimeMustBePast(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	roomdatalist := RoomDataList{
		Note:          "9999",
		Day:           2,
		EnterRoomTime: time.Now().Add(24 * time.Hour), // อนาคต, fail
	}

	ok, err := govalidator.ValidateStruct(roomdatalist)

	g.Expect(ok).ToNot(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("EnterRoomTime must be in the past"))
}
