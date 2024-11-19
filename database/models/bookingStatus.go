package models

type BookingStatus struct {
	Name         string `json:"name"`
	Descriptions string `json:"descriptions"`
	BaseModel
}
