package server

import (
	"fmt"
	"net/http"

	"github.com/learngo/rediscachehttpserver/src/handlers"
)

func RunServer() {
	http.HandleFunc("/get", handlers.GetHandler)
	http.HandleFunc("/set", handlers.SetHandler)
	http.HandleFunc("/", handlers.IndexHandler)
	serverAddress := "0.0.0.0:8000"
	fmt.Println(("Server Listening on\t:\nhttp://" + serverAddress))
	http.ListenAndServe(serverAddress, nil)
}
