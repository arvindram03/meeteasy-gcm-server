package main

import (
	"github.com/arvindram03/meeteasy-gcm-server/controllers"
	"github.com/bmizerany/pat"
	"io"
	"log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hey dude watsup ?!! \n")
}

func main() {
	m := pat.New()
	m.Get("/", http.HandlerFunc(HelloServer))

	m.Post("/user", http.HandlerFunc(controllers.RegisterUser))
	m.Post("/contacts/filter", http.HandlerFunc(controllers.GetRegisteredContacts))

	m.Post("/meetup", http.HandlerFunc(controllers.CreateMeetup))
	m.Put("/meetup/:meetupId", http.HandlerFunc(controllers.UpdateMeetup))

	m.Post("/message", http.HandlerFunc(controllers.SendMessage))

	http.Handle("/", m)
	log.Println("Starting MeetEasy GCM Server...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
