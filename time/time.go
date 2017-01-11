package time

import (
	"fmt"
	"time"
)

//　layout
const (
	MYSQL = "2006-01-02 15:04:05"
)

// mysql形式で文字列化
func StringMysql(t time.Time) string {
	return fmt.Sprintf("%v/%v/%v %v:%v:%v", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}
