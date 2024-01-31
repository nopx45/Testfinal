package entity

type Promotion struct {
	gorm.Model
	Firstname string `valid:"required~กรุณากรอกชื่อ !"`
	Lastname string
}