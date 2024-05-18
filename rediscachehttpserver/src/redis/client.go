package redis_client

import (
	"context"
	"log"
	"time"

	"github.com/learngo/rediscachehttpserver/src/utils"
	"github.com/redis/go-redis/v9"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:     utils.GetEnvOrDef("REDIS_ADDRESS", "localhost:6379"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

var ctx = context.Background()

func Set(key string, value string) {
	log.Println(RDB.Set(ctx, key, value, time.Minute))
}

func Get(key string) string {
	result, err := RDB.Get(ctx, key).Result()
	if err != nil {
		log.Output(1, err.Error())
		return "Not Available"
	}
	return result
}
