package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learngo/ginrediscurrenciesapi/src/services"
)

func returnJsonError(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func GetExchangeRatesHandler(c *gin.Context) {
	exchangeRates, err := services.GetExchangeRates()
	if err != nil {
		returnJsonError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": exchangeRates,
	})
}

func GetCurrenciesHandler(c *gin.Context) {
	currencies, err := services.GetCurrencies()
	if err != nil {
		returnJsonError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": currencies,
	})
}

func GetConversionHandler(c *gin.Context) {
	toCurrency, _ := c.Params.Get("toCurrency")
	fromCurrency, _ := c.Params.Get("fromCurrency")
	amountStr, _ := c.Params.Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		returnJsonError(c, err)
		return
	}
	value, err := services.GetConversion(
		toCurrency,
		fromCurrency,
		amount,
	)
	if err != nil {
		returnJsonError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": value,
	})
}
