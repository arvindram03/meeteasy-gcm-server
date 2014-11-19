package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"github.com/arvindram03/meeteasy-gcm-server/controllers"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hey dude watsup ?!! \n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/user",controllers.RegisterUser)
	http.HandleFunc("/meetup",controllers.CreateMeetup)
	log.Println("Starting MeetEasy GCM Server...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
