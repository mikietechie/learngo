package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, time.Now())
}
