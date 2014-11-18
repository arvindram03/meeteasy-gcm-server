package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"controllers"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hey dude watsup ?!! \n")
}

func main() {
	http.HandleFunc("/", HelloServer())
	http.HandleFunc("/user", controllers.Register)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
