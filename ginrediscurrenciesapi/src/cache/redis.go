package cache

import (
	"github.com/learngo/ginrediscurrenciesapi/src/utils"
	"github.com/redis/go-redis/v9"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:     utils.GetEnvOrDef("redisString", "localhost:6379"),
	Password: "", // no password set
	DB:       0,  // use default DB
})
