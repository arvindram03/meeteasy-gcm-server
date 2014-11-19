package dtos

type MeetupDTO struct {
	Location       LocationDTO `json:"location"`
	MobileNumber   string    `json:"mobileNumber"`
	Members        []UserDTO `json:"members"`
	Description    string    `json:"description"`
}

type LocationDTO struct {
	Latitude  float64	`json:"latitude"`
	Longitude float64	`json:"longitude"`
}
