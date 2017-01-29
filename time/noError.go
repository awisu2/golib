package time

import (
	"fmt"
)

// create YYYY/MM/DD HH:MM by mysql format strinig
func MysqlToMMDDHHMMNoError(text string) (date string) {
	date, err := MysqlToMMDDHHMM(text)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return
}

// create YYYY/MM/DD HH:MM by mysql format strinig
func MysqlToYYYYMMDDHHMMNoError(text string) (date string) {
	date, err := MysqlToYYYYMMDDHHMM(text)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return
}

// create YYYY/MM/DD by mysql format strinig
func MysqlToYYYYMMDDNoError(text string) (date string) {
	date, err := MysqlToYYYYMMDD(text)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return
}

// create MM/DD by mysql format strinig
func MysqlToMMDDNoError(text string) (date string) {
	date, err := MysqlToMMDD(text)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return
}

// create YYYY年MM月DD日 by mysql format strinig
func MysqlToYYYYMMDDJPNoError(text string) (date string) {
	date, err := MysqlToYYYYMMDDJP(text)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return
}

// create MM月DD日 by mysql format strinig
func MysqlToMMDDJPNoError(text string) (date string) {
	date, err := MysqlToMMDDJP(text)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return
}
