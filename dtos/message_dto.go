package dtos

type MessageDTO struct {
	SenderID          string      `json:"senderID"`
	Location          LocationDTO `json:"location"`
	TimeToReachInMins int         `json:"timeToReachInMins"`
	MeetupID          string      `json:"meetupID"`
}
