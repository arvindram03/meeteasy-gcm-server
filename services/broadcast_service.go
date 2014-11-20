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
	PASSWORD     = os.Getenv("GCM_API_KEY")
)

func Broadcast(message dtos.MessageDTO, registrationIDs []string) {
	broadcastMessage := serviceDTOs.GCMMessageDTO{
		Data:            message,
		RegistrationIDs: registrationIDs,
	}

	httpWrapper := &http.HttpWrapper{}
	request := httpWrapper.CreateRequest(GCM_SEND_URL, http.POST)
	request.SetBasicAuth(USERNAME, PASSWORD)
	request.SetJsonBody(broadcastMessage)
	_, err := request.Send()
	if err != nil {
		log.Println("[GCMError]: ", err)
	}
}
