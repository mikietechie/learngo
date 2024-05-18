package server

import (
	"fmt"
	"net/http"

	"github.com/learngo/dockeredgo/src/handlers"
)

func RunServer() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/time", handlers.TimeHandler)
	serverAddress := "0.0.0.0:8000"
	fmt.Println(("Server Listening on\t:\nhttp://" + serverAddress))
	http.ListenAndServe(serverAddress, nil)
}
