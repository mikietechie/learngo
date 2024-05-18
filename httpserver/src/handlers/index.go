package handlers

import (
	"net/http"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Here is a string...."))
}
