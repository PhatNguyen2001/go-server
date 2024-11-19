package models

type Booking struct {
	BookingDate      string        `json:"booking_date"`
	BookingHourStart string        `json:"booking_hour_start"`
	BookingHourEnd   string        `json:"booking_hour_end"`
	CustomerId       string        `json:"customer_id"`
	DoctorId         string        `json:"doctor_id"`
	ServiceId        string        `json:"service_id"`
	BookingStatusId  string        `json:"booking_status_id"`
	Customer         Customer      `gorm:"foreignKey:UserId"`
	Doctor           Doctor        `gorm:"foreignKey:DoctorId"`
	BookingStatus    BookingStatus `gorm:"foreignKey:BookingStatusId"`
	BaseModel
}
