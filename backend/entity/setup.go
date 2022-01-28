package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("roomlist.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Patient{},
		&Officer{},
		&Specialist{},
		&RoomDetail{},
		&RoomDataList{},
		&MedicineandVaccine{},
		&Category{},
		&Age{},
                &DosageForm{},
		&Room{},
		&Symptom{},
		&Screening{},
	)


	db = database

	//- Officer Data -
	password, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Mutant",
		Email:    "MUTANT@gmail.com",
		Password: string(password),
	})

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Titan",
		Email:    "TITAN@gmail.com",
		Password: string(password),
	})

	db.Model(&Officer{}).Create(&Officer{
		Name:     "เจ้าหน้าที่เอ",
		Email:    "aaa@example.com",
		Password: string(password),
	})

	db.Model(&Officer{}).Create(&Officer{
		Name:     "เจ้าหน้าที่บี",
		Email:    "bbb@example.com",
		Password: string(password),
	})

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Wittaya",
		Officer: "Wittaya",
		Email:    "wittaya@gmail.com",
		Password: string(password),
	})

	var Mutant Officer
	var Titan Officer
	var a Officer
	var b Officer
	var wittaya Officer
	db.Raw("SELECT * FROM officers WHERE email = ?", "MUTANT@gmail.com").Scan(&Mutant)
	db.Raw("SELECT * FROM officers WHERE email = ?", "TITAN@example.com").Scan(&Titan)
	db.Raw("Select * FROM officers WHERE email = ?", "aaa@example.com").Scan(&a)
	db.Raw("Select * FROM officers WHERE email = ?", "bbb@example.com").Scan(&b)
	db.Raw("SELECT * FROM officers WHERE email = ?", "wittaya@gmail.com").Scan(&wittaya)

	// - Patient Data -
	c := Patient{
		Name:     "นายซี",
		Behavior: "ไข้สูง ติดเชื้อ",
		Email:    "ccc@example.com",
		Tel:      "2222222222",
	}
	db.Model(&Patient{}).Create(&c)

	d := Patient{
		Name:     "นายดี",
		Behavior: "หวัดรุนเเรง ติดเชื้อทางเดินหายใจ",
		Email:    "ddd@example.com",
		Tel:      "3333333333",
	}
	db.Model(&Patient{}).Create(&d)

	e := Patient{
		Name:     "นายอี",
		Behavior: "ติดเชื้อทางผิวหนัง มีเชื้อที่แผลตามตัว",
		Email:    "fff@example.com",
		Tel:      "4444444444",
	}
	db.Model(&Patient{}).Create(&e)

	patient1 := Patient{
		Name: "aaaaa",
		Age: 28,
		Gender: "Female",
		BloodPressure: 150,
	}
	db.Model(&Patient{}).Create(&patient1)

	patient2 := Patient{
		Name: "bbbbb",
		Age: 13,
		Gender: "Male",
		BloodPressure: 170,
	}
	db.Model(&Patient{}).Create(&patient2)

	// - room data -

	room1 := Room{
		RoomNumber: "b201",
	}
	db.Model(&Room{}).Create(&room1)

	room2 := Room{
		RoomNumber: "b202",
	}
	db.Model(&Room{}).Create(&room2)

	// - symptom data -

	symptom1 := Symptom{
		State: "head ache",
		Period: 3,
	}
	db.Model(&Symptom{}).Create(&symptom1)

	symptom2 := Symptom{
		State: "cough",
		Period: 5,
	}
	db.Model(&Symptom{}).Create(&symptom2)
	

	// - Specialist Data -
	u := Specialist{
		Name:  "ดร.หมอยู",
		Email: "FFF@example.com",
	}
	db.Model(&Specialist{}).Create(&u)

	g := Specialist{
		Name:  "ดร.หมอจี",
		Email: "GGG@example.com",
	}
	db.Model(&Specialist{}).Create(&g)

	h := Specialist{
		Name:  "ดร.หมอฮง",
		Email: "HHH@example.com",
	}
	db.Model(&Specialist{}).Create(&h)

	// - RoomDetail Data -
	Room01 := RoomDetail{
		Name: "1001",
		Size: "VIP large",
		Cost: "6500 Baht/Day",
	}
	db.Model(&RoomDetail{}).Create(&Room01)

	Room02 := RoomDetail{
		Name: "1002",
		Size: "VIP medium",
		Cost: "5500 Baht/Day",
	}
	db.Model(&RoomDetail{}).Create(&Room02)

	Room03 := RoomDetail{
		Name: "1003",
		Size: "VIP small",
		Cost: "4000 Baht/Day",
	}
	db.Model(&RoomDetail{}).Create(&Room03)

	Room04 := RoomDetail{
		Name: "1004",
		Size: "large",
		Cost: "3500 Baht/Day",
	}
	db.Model(&RoomDetail{}).Create(&Room04)

	Room05 := RoomDetail{
		Name: "1005",
		Size: "medium",
		Cost: "2500 Baht/Day",
	}
	db.Model(&RoomDetail{}).Create(&Room05)

	Room06 := RoomDetail{
		Name: "1006",
		Size: "small",
		Cost: "1500 Baht/Day",
	}
	db.Model(&RoomDetail{}).Create(&Room06)

	// -RoomDataList-

	db.Model(&RoomDataList{}).Create(&RoomDataList{
		Room:          Room05,
		Officer:       a,
		Patient:       d,
		Specialist:    h,
		Day:           3,
		Note:          "ห้ามคนนอกเข้า ก่อนได้รับชุดป้องกัน ห้ามสัมผัสผู้ป่วยโดยตรง",
		EnterRoomTime: time.Now(),
	})

	db.Model(&RoomDataList{}).Create(&RoomDataList{
		Room:          Room02,
		Officer:       Mutant,
		Patient:       c,
		Specialist:    g,
		Day:           6,
		Note:          "ห้ามสัมผัสผู้ป่วยโดยตรง",
		EnterRoomTime: time.Now(),
	})

	db.Model(&RoomDataList{}).Create(&RoomDataList{
		Room:          Room03,
		Officer:       b,
		Patient:       e,
		Specialist:    u,
		Day:           4,
		Note:          "ห้ามสัมผัสผู้ป่วยโดยตรง",
		EnterRoomTime: time.Now(),
	})
	
		//Category data
	Medicine := Category{
		Category: "ยา",
	}
	db.Model(&Category{}).Create(&Medicine)

	Vaccine := Category{
		Category: "วัคซีน",
	}
	db.Model(&Category{}).Create(&Vaccine)

	//DosageForm data
	None := DosageForm{
		DosageForm: "-",
	}

	db.Model(&DosageForm{}).Create(&None)
	Water := DosageForm{
		DosageForm: "ยาน้ำ",
	}
	db.Model(&DosageForm{}).Create(&Water)

	Capsule := DosageForm{
		DosageForm: "ยาแคปซูล",
	}
	db.Model(&DosageForm{}).Create(&Capsule)

	Tablet := DosageForm{
		DosageForm: "ยาเม็ด",
	}
	db.Model(&DosageForm{}).Create(&Tablet)

	Perenteral := DosageForm{
		DosageForm: "ยาฉีด",
	}
	db.Model(&DosageForm{}).Create(&Perenteral)

	//Age data
	MY := Age{
		Age: "เดือน-ปี",
	}
	db.Model(&Age{}).Create(&MY)

	YY := Age{
		Age: "ปี-ปี",
	}
	db.Model(&Age{}).Create(&YY)
}	
