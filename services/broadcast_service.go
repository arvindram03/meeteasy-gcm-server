package services

import (
	"github.com/arvindram03/meeteasy-gcm-server/dtos"
	"github.com/arvindram03/meeteasy-gcm-server/http"
	serviceDTOs "github.com/arvindram03/meeteasy-gcm-server/services/dtos"
	"log"
	"os"
)

const (
	GCM_SEND_URL = "https://android.googleapis.com/gcm/send"
	USERNAME     = "key"
)

func Broadcast(message dtos.MessageDTO, registrationIDs []string) {
	password  := os.Getenv("GCM_API_KEY")
	broadcastMessage := serviceDTOs.GCMMessageDTO{
		Data:            message,
		RegistrationIDs: registrationIDs,
	}

	httpWrapper := &http.HttpWrapper{}
	request := httpWrapper.CreateRequest(GCM_SEND_URL, http.POST)
	request.SetBasicAuth(USERNAME, password)
	request.SetJsonBody(broadcastMessage)
	_, err := request.Send()
	if err != nil {
		log.Println("[GCMError]: ", err)
	}
}
