package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/learngo/ginrediscurrenciesapi/src/cache"
	"github.com/learngo/ginrediscurrenciesapi/src/config"
	"github.com/learngo/ginrediscurrenciesapi/src/structs"
)

func GetExchangeRates() (structs.BeaconResponse, error) {
	logger := log.New(os.Stdout, "Get Exchange Rates ", log.Flags())
	key := "ExchangeRates"
	var obj structs.BeaconResponse
	data, _ := cache.RDB.Get(config.CTX, key).Result()
	if data != "" {
		logger.Println("In Cache")
		json.Unmarshal([]byte(data), &obj)
		return obj, nil
	}
	logger.Println("Not In Cache")
	url := fmt.Sprintf("%s/latest/?base=%s&api_key=%s", config.BEACON_URL, config.BEACON_BASE_CURRENCY, config.BEACON_KEY)
	res, err := http.Get(url)
	if err != nil {
		return obj, err
	}
	if res.StatusCode != 200 {
		return obj, errors.New("server returned something else")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return obj, err
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return obj, err
	}
	cache.RDB.Set(config.CTX, key, string(body), time.Minute*15)
	logger.Println("Now In Cache")
	return obj, nil
}

func GetCurrencies() ([]string, error) {
	obj, err := GetExchangeRates()
	if err != nil {
		return nil, err
	}
	var currencies []string
	for k := range obj.Data.Rates {
		currencies = append(currencies, k)
	}
	return currencies, nil
}

func GetConversion(toCurrency, fromCurrency string, amount float64) (float64, error) {
	obj, err := GetExchangeRates()
	if err != nil {
		return 0, err
	}
	var toRate, fromRate float64 = 1, 1
	if toCurrency != config.BEACON_BASE_CURRENCY {
		toRate = obj.Data.Rates[toCurrency]
		if toRate == 0 {
			return 0, nil
		}
	}
	if fromCurrency != config.BEACON_BASE_CURRENCY {
		fromRate = obj.Data.Rates[fromCurrency]
		if fromRate == 0 {
			return 0, nil
		}
	}
	return amount * toRate / fromRate, nil
}
