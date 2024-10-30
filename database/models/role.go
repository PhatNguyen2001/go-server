package models

type Role struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Setting     string `json:"setting"`
	Description string `json:"password"`
	*BaseModel
}
