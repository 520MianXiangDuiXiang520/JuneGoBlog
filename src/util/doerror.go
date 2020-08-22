package util

import (
	"log"
	"runtime"
)

func CatchException(e error) {
	if e != nil {
		pc, file, line, _ := runtime.Caller(1)
		fName := runtime.FuncForPC(pc).Name()
		log.Printf("Error: %v(%v): || [%v] 执行异常：%v", file, line, fName, e)
		panic(e)
	}
}
