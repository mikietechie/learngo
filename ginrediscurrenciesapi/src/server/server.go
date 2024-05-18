package server

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/learngo/ginrediscurrenciesapi/src/config"
	"github.com/learngo/ginrediscurrenciesapi/src/handlers"
	"github.com/learngo/ginrediscurrenciesapi/src/middleware"
)

func RunServer() {
	server := gin.Default()
	server.Use(gin.Logger())
	server.Use(cors.Default())
	server.Use(gin.Recovery())
	server.Use(middleware.Security())
	// Routing
	server.GET("/exchange-rates", handlers.GetExchangeRatesHandler)
	server.GET("/currencies", handlers.GetCurrenciesHandler)
	server.GET("/conversion/:toCurrency/:fromCurrency/:amount", handlers.GetConversionHandler)
	server.GET("/", handlers.IndexHandler)
	// End of routing

	// Run Server
	log.Println(("Server Listening on\t:\nhttp://" + config.SERVER_ADDRESS))
	http.ListenAndServe(config.SERVER_ADDRESS, server)
}
