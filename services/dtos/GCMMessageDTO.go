package dtos

type GCMMessageDTO struct {
	Data            interface{} `json:"data"`
	RegistrationIDs []string    `json:"registration_ids"`
}
