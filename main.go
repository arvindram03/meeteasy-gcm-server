package main

import (
	"io"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":6055", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
