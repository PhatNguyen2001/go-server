package models

type Manager struct {
	Gender    string `json:"gender"`
	Bod       string `json:"bod"`
	AvatarUrl string `json:"avatar_url"`
	UserId    string `json:"user_id"`
	User      User   `gorm:"foreignKey:UserId"`
	BaseModel
}
