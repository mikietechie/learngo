package handlers

import (
	"fmt"
	"net/http"

	redis_client "github.com/learngo/rediscachehttpserver/src/redis"
)

func SetHandler(res http.ResponseWriter, req *http.Request) {
	redis_client.Set(
		req.URL.Query().Get("key"),
		req.URL.Query().Get("value"),
	)
	fmt.Fprint(res, req.URL.Query().Get("key"))
}

func GetHandler(res http.ResponseWriter, req *http.Request) {
	value := redis_client.Get(req.URL.Query().Get("key"))
	fmt.Fprint(res, value)
}
