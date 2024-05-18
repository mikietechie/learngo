package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/learngo/ginrediscurrenciesapi/src/utils"
)

var _ = godotenv.Overload()

var CTX = context.Background()

var SERVER_ADDRESS = utils.GetEnvOrDef("SERVER_ADDRESS", "0.0.0.0:8000")
var REDIS_CONNECTION = utils.GetEnvOrDef("REDIS_CONNECTION", "localhost:6379")

var BEACON_KEY = utils.GetEnvOrDef("BEACON_KEY", "")
var BEACON_URL = utils.GetEnvOrDef("BEACON_URL", "https://api.currencybeacon.com/v1")
var BEACON_BASE_CURRENCY = utils.GetEnvOrDef("BEACON_BASE_CURRENCY", "USD")
