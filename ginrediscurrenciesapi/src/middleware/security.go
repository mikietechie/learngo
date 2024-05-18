package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.Header)
		// token := c.Request.Header.Get("Authorization")
		// TODO := check if user is in mongo db
		c.Next()
	}
}
