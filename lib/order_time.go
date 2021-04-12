package lib

import (
	"fmt"
	"time"
)

func GetOrderIdTime() string {
	currentTime := time.Now().Nanosecond()
	return fmt.Sprintf("%d", currentTime)
}
