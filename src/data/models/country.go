package models

type Country struct {
	BaseModel
	Name   string `gorm:"size15;type:string;not null"`
	Cities *[]City
}
