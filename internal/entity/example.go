package entity

type Example struct {
	ID   uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"name" json:"name"`
}
