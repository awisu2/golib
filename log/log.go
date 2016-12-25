package log

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

const (
	CALLER_START        = 3
	CALLER_LENGTH       = 1
	CALLER_LENGTH_ERROR = 4
)

// override log.Println
func Errorf(v ...interface{}) (err error) {
	caller := getCallers(CALLER_LENGTH_ERROR)

	s := fmt.Sprint(v...)
	log.Println("Error : " + s + "\n" + caller)

	err = fmt.Errorf("%v", s)
	return
}

// override log.Println
func Println(v ...interface{}) {
	log.Println(fmt.Sprint(v...), getCallers(CALLER_LENGTH))
}

// override log.Panicln
func Panicln(v ...interface{}) {
	log.Panicln(fmt.Sprint(v...), "\n", getCallers(CALLER_LENGTH_ERROR))
}

// override log.Fatalln
func Fatalln(v ...interface{}) {
	log.Fatalln(fmt.Sprint(v...), getCallers(CALLER_LENGTH_ERROR))
}

func getCaller(num int) string {
	s := ""
	pc, file, line, ok := runtime.Caller(num)
	if ok {
		s = " " + fmt.Sprintf("[%v]%v:%v", pc, file, line)
	}
	return s
}

func getCallers(length int) (s string) {

	for i := CALLER_START; i < CALLER_START+length; i++ {
		_s := getCaller(i)
		if _s == "" {
			break
		}
		s += "\n" + strconv.Itoa(i) + " : " + _s
	}
	if len(s) > 0 {
		s = s[1:]
	}
	return
}
