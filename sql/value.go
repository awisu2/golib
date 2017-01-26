package sql

import (
	"fmt"
	"time"
)

func Now() string {
	t := time.Now()
	date := fmt.Sprintf("%d/%d/%d %d:%d:%d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())

	return date
}
