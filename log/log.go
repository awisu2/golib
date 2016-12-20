package log

import (
	"fmt"
	"log"
	"runtime"
)

// override log.Println
func Errorf(v ...interface{}) (err error) {
	caller := getCaller()

	s := fmt.Sprint(v...)
	log.Println("Error : " + s + " " + caller)

	err = fmt.Errorf("%v", s)
	return
}

// override log.Println
func Println(v ...interface{}) {
	log.Println(fmt.Sprint(v...), getCaller())
}

// override log.Panicln
func Panicln(v ...interface{}) {
	log.Panicln(fmt.Sprint(v...), getCaller())
}

// override log.Fatalln
func Fatalln(v ...interface{}) {
	log.Fatalln(fmt.Sprint(v...), getCaller())
}

func getCaller() string {
	s := ""
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		s = " " + fmt.Sprintf("[%v]%v:%v", pc, file, line)
	}
	return s
}
