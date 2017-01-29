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

// parse mysql format
func ParseMySql(text string) (t time.Time, err error) {
	return time.Parse(MYSQL, text)
}

// parce and convert format
func ConvertMysql(text string, f func(t time.Time) string) (date string, err error) {
	t, err := ParseMySql(text)
	if err != nil {
		return
	}
	return f(t), nil
}

// create YYYY/MM/DD HH:MM by mysql format strinig
func MysqlToYYYYMMDDHHMM(text string) (date string, err error) {
	return ConvertMysql(text, func(t time.Time) string {
		return fmt.Sprintf("%v/%v/%v %v:%v", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())
	})
}

// create MM/DD HH:MM by mysql format strinig
func MysqlToMMDDHHMM(text string) (date string, err error) {
	return ConvertMysql(text, func(t time.Time) string {
		return fmt.Sprintf("%v/%v %v:%v", int(t.Month()), t.Day(), t.Hour(), t.Minute())
	})
}

// create YYYY/MM/DD by mysql format strinig
func MysqlToYYYYMMDD(text string) (date string, err error) {
	return ConvertMysql(text, func(t time.Time) string {
		return fmt.Sprintf("%v/%v/%v", t.Year(), int(t.Month()), t.Day())
	})
}

// create MM/DD by mysql format strinig
func MysqlToMMDD(text string) (date string, err error) {
	return ConvertMysql(text, func(t time.Time) string {
		return fmt.Sprintf("%v/%v", int(t.Month()), t.Day())
	})
}

// create YYYY年MM月DD日 by mysql format strinig
func MysqlToYYYYMMDDJP(text string) (date string, err error) {
	return ConvertMysql(text, func(t time.Time) string {
		return fmt.Sprintf("%v年%v月%v日", t.Year(), int(t.Month()), t.Day())
	})
}

// create MM月DD日 by mysql format strinig
func MysqlToMMDDJP(text string) (date string, err error) {
	return ConvertMysql(text, func(t time.Time) string {
		return fmt.Sprintf("%v月%v日", int(t.Month()), t.Day())
	})
}
