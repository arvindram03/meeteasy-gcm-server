package dtos

type MeetupDTO struct {
	MeetupTime        string `json:"meetupTime"`
	Location          LocationDTO `json:"location"`
	MobileNumber      string    `json:"mobileNumber"`
	MemberIDs           []string `json:"memberIDs"`
	Description       string    `json:"description"`
}

type LocationDTO struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
}
