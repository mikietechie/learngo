package random

import "math/rand"

func GetRandomInt(max int) int64 {
	return int64(rand.Intn(max))
}
