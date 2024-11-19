package models

type BookingDetail struct {
	OriginalPrice    string `json:"original_price"`
	PromotionalPrice string `json:"promotional_price"`
	PaymentPrice     string `json:"payment_price"`
	BookingId        string `json:"booking_id"`

	Booking Booking `gorm:"foreignKey:BookingId"`
	BaseModel
}
