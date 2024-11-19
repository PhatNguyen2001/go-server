package models

type Doctor struct {
	Gender        string `json:"gender"`
	Bod           string `json:"bod"`
	AvatarUrl     string `json:"avatar_url"`
	Level         string `json:"level"`
	TotalBookings string `json:"total_bookings"`
	UserId        string `json:"user_id"`
	User          User   `gorm:"foreignKey:UserId"`
	BaseModel
}
