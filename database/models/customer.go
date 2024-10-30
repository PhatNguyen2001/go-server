package models

type Cusomter struct {
	Gender        string `gorm:"primaryKey" json:"id"`
	Bod           string `json:"name"`
	AvatarUrl     string `json:"setting"`
	Level         string `json:"password"`
	TotalBookings string `json:"total_bookings"`
	TotalExpenses string `json:"total_expenses"`
	UserId        string `json:"user_id"`
	User          User   `gorm:"foreignKey:UserId"`
	BaseModel
}
