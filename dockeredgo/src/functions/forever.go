package functions

import "time"

func RunLogTimeForever() {
	for {
		LogTime()
		time.Sleep(time.Second)
	}
}
