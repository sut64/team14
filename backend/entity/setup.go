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
	database, err := gorm.Open(sqlite.Open("team14.db"), &gorm.Config{})
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
		&Appointment{},
		&Germ{},
		&CatchingType{},
		&RiskGroupType{},
		&Contagious{},
		&Prevention{},
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
		Email:    "wittaya@gmail.com",
		Password: string(password),
	})

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Jirawan",
		Email:    "jirawan@gmail.com",
		Password: string(password),
	})

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Nukanda",
		Email:    "nukanda@gmail.com",
		Password: string(password),
	})

	var Mutant Officer
	var Titan Officer
	var a Officer
	var b Officer
	var wittaya Officer
	var jirawan Officer
	var nukanda Officer

	db.Raw("SELECT * FROM officers WHERE email = ?", "MUTANT@gmail.com").Scan(&Mutant)
	db.Raw("SELECT * FROM officers WHERE email = ?", "TITAN@example.com").Scan(&Titan)
	db.Raw("Select * FROM officers WHERE email = ?", "aaa@example.com").Scan(&a)
	db.Raw("Select * FROM officers WHERE email = ?", "bbb@example.com").Scan(&b)
	db.Raw("SELECT * FROM officers WHERE email = ?", "wittaya@gmail.com").Scan(&wittaya)
	db.Raw("SELECT * FROM officers WHERE email = ?", "jirawan@gmail.com").Scan(&jirawan)
	db.Raw("SELECT * FROM officers WHERE email = ?", "nukanda@gmail.com").Scan(&nukanda)

	// - Patient Data -
	c := Patient{
		Name:          "นายซี",
		Age:           17,
		Gender:        "Male",
		BloodPressure: 170,
	}
	db.Model(&Patient{}).Create(&c)

	d := Patient{
		Name:          "นายดี",
		Age:           18,
		Gender:        "Male",
		BloodPressure: 170,
	}
	db.Model(&Patient{}).Create(&d)

	e := Patient{
		Name:          "นายอี",
		Age:           19,
		Gender:        "Male",
		BloodPressure: 170,
	}
	db.Model(&Patient{}).Create(&e)

	patient1 := Patient{
		Name:          "aaaaa",
		Age:           28,
		Gender:        "Female",
		BloodPressure: 150,
	}
	db.Model(&Patient{}).Create(&patient1)

	patient2 := Patient{
		Name:          "bbbbb",
		Age:           13,
		Gender:        "Male",
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
		State:  "head ache",
		Period: 3,
	}
	db.Model(&Symptom{}).Create(&symptom1)

	symptom2 := Symptom{
		State:  "cough",
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
	MM := Age{
		Age: "เดือน - เดือน",
	}
	db.Model(&Age{}).Create(&MM)
	MY := Age{
		Age: "เดือน - ปี",
	}
	db.Model(&Age{}).Create(&MY)

	YY := Age{
		Age: "ปี - ปี",
	}
	db.Model(&Age{}).Create(&YY)

	// Germ Data
	g1 := Germ{
		Name: "ไข้หวัดใหญ่สายพันธ์ุเอ",
	}
	db.Model(&Germ{}).Create(&g1)

	g2 := Germ{
		Name: "เดงกี",
	}
	db.Model(&Germ{}).Create(&g2)

	g3 := Germ{
		Name: "ลาสซา",
	}
	db.Model(&Germ{}).Create(&g3)

	g4 := Germ{
		Name: "วาริโอลา",
	}
	db.Model(&Germ{}).Create(&g4)

	g5 := Germ{
		Name: "โคโรนาสายพันธุ์ใหม่ 2019",
	}
	db.Model(&Germ{}).Create(&g5)

	g6 := Germ{
		Name: "ไข้หวัดใหญ่สายพันธุ์บี",
	}
	db.Model(&Germ{}).Create(&g6)

	// CatchingType Data
	animal_to_human := CatchingType{
		Title: "ติดต่อจากสัตว์สู่คน",
	}
	db.Model(&CatchingType{}).Create(&animal_to_human)

	human_to_human := CatchingType{
		Title: "ติดต่อจากคนสู่คน",
	}
	db.Model(&CatchingType{}).Create(&human_to_human)

	animal_to_animal := CatchingType{
		Title: "ติดต่อจากสัตว์สู่สัตว์",
	}
	db.Model(&CatchingType{}).Create(&animal_to_animal)

	human_to_animal := CatchingType{
		Title: "ติดต่อจากคนสู่สัตว์",
	}
	db.Model(&CatchingType{}).Create(&human_to_animal)

	// RiskGroupType Data
	RG1 := RiskGroupType{
		Title: "เด็กเล็ก",
	}
	db.Model(&RiskGroupType{}).Create(&RG1)

	RG2 := RiskGroupType{
		Title: "ผู้สูงอายุ",
	}
	db.Model(&RiskGroupType{}).Create(&RG2)

	RG3 := RiskGroupType{
		Title: "บุคคลทั่วไป",
	}
	db.Model(&RiskGroupType{}).Create(&RG3)

	// Contagious Data
	C1 := Contagious{
		Name:          "ไข้ลาสซา",
		Germ:          g3,
		CatchingType:  animal_to_human,
		Symptom:       "มีไข้ อ่อนแรง ไม่สบายตัว มีเลือดออกง่ายผิดปกติ",
		Incubation:    14,
		RiskGroupType: RG3,
		Date:          time.Now().AddDate(-10, -5, -10),
	}
	db.Model(&Contagious{}).Create(&C1)

	C2 := Contagious{
		Name:          "ไข้เลือดออก",
		Germ:          g2,
		CatchingType:  animal_to_human,
		Symptom:       "มีไข้ ปวดหัว วิงเวียน คลื่นไส้อาเจียน ปวดรอบกระบอกตา ปวดตามกล้ามเนื้อ",
		Incubation:    7,
		RiskGroupType: RG3,
		Date:          time.Now().AddDate(-5, -10, -10),
	}
	db.Model(&Contagious{}).Create(&C2)

	C3 := Contagious{
		Name:          "โรคฝีดาษ",
		Germ:          g4,
		CatchingType:  human_to_human,
		Symptom:       "มีไข้สูง รู้สึกไม่สบายตัว หนาวสั่น ปวดศีรษะ อ่อนเพลียอย่างรุนแรง ปวดหลังอย่างรุนแรง อาเจียน",
		Incubation:    12,
		RiskGroupType: RG3,
		Date:          time.Now().AddDate(-2, -10, -5),
	}
	db.Model(&Contagious{}).Create(&C3)

	// Prevention Data
	db.Model(&Prevention{}).Create(&Prevention{
		Officer:    Mutant,
		Contagious: C1,
		Disease:    "เป็นชนิดของไข้เลือดออกจากไวรัสที่เกิดจากไวรัสลาสซา ไวรัสในวงศ์ Arenaviridae",
		Specialist: g,
		Protection: "ใช้การแยกผู้ติดเชื้อและลดการสัมผัสกับหนู",
		Date:       time.Now(),
		Age:        30,
	})

	db.Model(&Prevention{}).Create(&Prevention{
		Officer:    b,
		Contagious: C2,
		Disease:    "มียุงลายเป็นพาหะนำโรค มักพบในประเทศเขตร้อนและระบาดในช่วงฤดูฝนของทุกปี",
		Specialist: u,
		Protection: "ป้องกันไม่ให้ยุงลายกัด โดยสวมใส่เสื้อผ้าที่ปกปิดมิดชิด ใช้สารไล่ยุงชนิดต่างๆ เช่น DEET รวมถึงป้องกันไม่ให้ยุงลายเข้ามาหลบซ่อนในบ้าน ทั้งนี้ ยุงลายมักกัดในเวลากลางวันมากกว่ากลางคืน",
		Date:       time.Now(),
		Age:        32,
	})
}
